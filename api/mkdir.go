package api

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tarocch1/file-admin/common"
)

func MkdirHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	// 获取 path 参数
	path := r.FormValue("path")
	dir := r.FormValue("dir")

	// 计算出工作路径
	workingPath, err := common.GetWorkingPath(path)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	targetPath := filepath.Join(workingPath, dir)
	if !common.PathNotExist(targetPath) {
		ErrorHandler(w, http.StatusConflict, errors.New("dir exists"))
		return
	}

	err = os.MkdirAll(targetPath, 0755)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, err)
		return
	}

	JsonHandler(w, make(map[string]interface{}))
}
