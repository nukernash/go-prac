package file_processor

import (
	"container/list"
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileProcessor struct {
	Root string
}

type File struct {
	contentHash string
	name        string
}

type Dir struct {
	files []string
	dirs  []string
}

func (f *FileProcessor) ProcessConcurrently() [][]string {

	fileMap := make(map[string][]string)
	dirChan := make(chan string)
	dirDataChan := make(chan Dir)

	fileChan := make(chan string)
	fileDataChan := make(chan File)

	doneChan := make(chan string)

	var dirWg sync.WaitGroup
	var fileWg sync.WaitGroup

	workerCount := 3
	for range workerCount {
		go f.directoryProcessing(dirChan, dirDataChan, &dirWg)
	}

	for range workerCount {
		go f.fileProcessing(fileChan, fileDataChan, &fileWg)
	}

	dirWg.Add(1)
	dirChan <- f.Root

	go func() {
		for {
			select {
			case dirData := <-dirDataChan:
				for _, dir := range dirData.dirs {
					dirWg.Add(1)
					dirChan <- dir
				}
				for _, file := range dirData.files {
					fileWg.Add(1)
					fileChan <- file
				}
			case fileData := <-fileDataChan:
				if l, exist := fileMap[fileData.contentHash]; exist {
					l = append(l, fileData.name)
					fileMap[fileData.contentHash] = l
				} else {
					fileMap[fileData.contentHash] = []string{fileData.name}
				}
			case <-doneChan:
				log.Println("All tasks ara done")
				close(dirChan)
				close(dirDataChan)
				close(fileChan)
				close(fileDataChan)
				return
			default:
				log.Print("")
				time.Sleep(time.Millisecond)
			}
		}
	}()

	dirWg.Wait()
	fileWg.Wait()

	doneChan <- "done"

	result := make([][]string, 0)
	for _, v := range fileMap {
		result = append(result, v)
	}
	return result
}

func (f *FileProcessor) directoryProcessing(dirChan <-chan string, dirDataChan chan<- Dir, wg *sync.WaitGroup) {
	for dir := range dirChan {
		files, dirs := f.getFileAndDirs(dir)
		dirData := &Dir{
			files: files,
			dirs:  dirs,
		}
		dirDataChan <- *dirData
		wg.Done()
	}
}

func (f *FileProcessor) fileProcessing(fileChan <-chan string, fileDataChan chan<- File, wg *sync.WaitGroup) {
	for file := range fileChan {
		hash := f.hash(f.getFileContent(file))
		fileData := File{
			contentHash: hash,
			name:        file,
		}

		fileDataChan <- fileData
		wg.Done()
	}
}

func (f *FileProcessor) ProcessSequentially() [][]string {
	result := make([][]string, 0)

	fileMap := make(map[string][]string)
	queue := list.New()
	queue.PushBack(f.Root)

	for queue.Len() != 0 {
		curr := queue.Front()
		item, ok := curr.Value.(string)
		if !ok {
			log.Printf("error while getting item from a list")
			queue.Remove(curr)
			continue
		}
		files, dirs := f.getFileAndDirs(item)
		for _, dir := range dirs {
			queue.PushBack(dir)
		}
		for _, file := range files {
			hash := f.hash(f.getFileContent(file))
			if l, exist := fileMap[hash]; exist {
				l = append(l, file)
				fileMap[hash] = l
			} else {
				fileMap[hash] = []string{file}
			}
		}
		queue.Remove(curr)
	}

	for _, v := range fileMap {
		result = append(result, v)
	}

	return result
}

func (f *FileProcessor) getFileAndDirs(dirName string) (files []string, dirs []string) {

	files = make([]string, 0)
	dirs = make([]string, 0)

	items, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Printf("error while listing item in a dir %s\n : %+v", dirName, err)
		return
	}

	for _, item := range items {
		if item.IsDir() {
			dirs = append(dirs, filepath.Join(dirName, item.Name()))
		} else if isFile(item) {
			files = append(files, filepath.Join(dirName, item.Name()))
		}
	}

	return
}

func (f *FileProcessor) getFileContent(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error while reading file %s\n : %+v", fileName, err)
		return ""
	}

	return string(content)
}

func (f *FileProcessor) hash(content string) string {
	hash := md5.Sum([]byte(content))
	return fmt.Sprintf("%x", hash)
}

func isSymlink(entry os.DirEntry) bool {
	return entry.Type()&os.ModeSymlink != 0
}

func isFile(entry os.DirEntry) bool {
	return entry.Type().IsRegular()
}
