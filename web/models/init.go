package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	log "github.com/cihub/seelog"
)

func Init() {
	task_manager = &manager{}
	task_manager.readin_task_list()
	sort.SliceStable(task_manager.task_list, func(x, y int) bool {
		a := task_manager.task_list[x]
		b := task_manager.task_list[y]
		return a.create.Unix() < b.create.Unix()
	})
}

func GetTaskManager() TaskManager {
	return task_manager
}

func (m *manager) readin_task_list() {

	filepath.Walk(Task_data_dir, func(p string, info os.FileInfo, e error) error {

		if e != nil {
			return e
		}

		if info.IsDir() && p != Task_data_dir {
			m.readin_task(p)
		}

		return nil
	})
}

func (m *manager) readin_task(dir string) error {

	idx := strings.LastIndex(dir, string(os.PathSeparator))
	id := dir[idx+1:]

	f := dir + "/status.json"

	cont, e := ioutil.ReadFile(f)

	if e != nil {
		log.Errorf("failed to read task status file %v: %v", f, e)
		return e
	}

	var status task_status

	if e = json.Unmarshal(cont, &status); e != nil {
		log.Errorf("failed to parse task's status file %v: %v", f, e)
		return e
	}

	item := &task_item{
		id:     id,
		create: status.Create,
		status: status.Status,
	}

	m.task_list = append(m.task_list, item)

	return nil
}