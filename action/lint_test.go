package action

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLintSuccess(t *testing.T) {
	tmpHome := createTmpHome()
	fakeUpdate(tmpHome)

	chartName := "goodChart"

	Create(chartName, tmpHome)

	output := capture(func() {
		Lint(chartName, tmpHome)
	})

	expected := "Chart [goodChart] has passed all necessary checks"

	if !strings.Contains(output, expected) {
		t.Fatalf("Expected: '%s' in %s.", expected, output)
	}
}

func TestLintMissingReadme(t *testing.T) {
	tmpHome := createTmpHome()
	fakeUpdate(tmpHome)

	chartName := "badChart"

	Create(chartName, tmpHome)

	os.Remove(filepath.Join(tmpHome, WorkspaceChartPath, chartName, "README.md"))

	output := capture(func() {
		Lint(chartName, tmpHome)
	})

	expected := "A README file was not found"

	if !strings.Contains(output, expected) {
		t.Fatalf("Expected: '%s' in %s.", expected, output)
	}
}

func TestLintAll(t *testing.T) {
	tmpHome := createTmpHome()
	fakeUpdate(tmpHome)

	missingReadmeChart := "missingReadme"

	Create(missingReadmeChart, tmpHome)
	os.Remove(filepath.Join(tmpHome, WorkspaceChartPath, missingReadmeChart, "README.md"))

	Create("goodChart", tmpHome)

	output := capture(func() {
		Cli().Run([]string{"helm", "--home", tmpHome, "lint", "--all"})
	})

	expectMatches(t, output, "A README file was not found.*"+missingReadmeChart)
	expectContains(t, output, "Chart [goodChart] has passed all necessary checks")
	expectContains(t, output, "Chart [missingReadme] failed some checks")
}
