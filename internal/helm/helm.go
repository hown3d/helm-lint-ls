package helm

import (
	"fmt"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/mrjosh/helm-lint-ls/internal/log"
	"helm.sh/helm/v3/pkg/chartutil"
)

var logger = log.GetLogger()

func IsChartDirectory(rootDir string) (bool, error) {
	return chartutil.IsChartDir(rootDir)
}

var globalValues chartutil.Values

func lookupValue(name string) (chartutil.Values, error) {
	return globalValues.Table(name)
}

func InitializeValues(chartDir string) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	// start synchronization for values.yaml file
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					dir := filepath.Dir(event.Name)
					v, err := refreshValues(dir)
					if err != nil {
						logger.Println(err)
					}
					globalValues = v
				}
			}
		}
	}()

	err = watcher.Add(filepath.Join(chartDir, "values.yaml"))
	if err != nil {
		return fmt.Errorf("adding values.yaml to fsnotify watcher: %w", err)
	}
	v, err := refreshValues(chartDir)
	globalValues = v
	return err
}

func refreshValues(chartDir string) (chartutil.Values, error) {
	valuesFile := filepath.Join(chartDir, "values.yaml")
	return chartutil.ReadValuesFile(valuesFile)
}
