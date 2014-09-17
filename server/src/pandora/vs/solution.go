package vs
import (
    "io/ioutil"
    "strings"
    "errors"
    "regexp"
)

type VsSolution struct {
    Name string
    Guid string
    Projects []*VsProject
    FilePath string
};

func getMatchesAsMap (r *regexp.Regexp, s string) (map[string]string) {
    matches := r.FindAllStringSubmatch (s, -1)[0]
    names := projectRegexId.SubexpNames()
    md := map[string]string{}
    for i, n := range matches {
         md[names[i]] = n
    }
    return md
}

var projectRegexId = regexp.MustCompile (`^Project\("(?P<guid>[^"]+)"\) = "(?P<name>[^"]+)", "(?P<project_path>[^"]+)", "(?P<project_guid>[^"]+)"`)
func parseSolution (filePath string) (solution *VsSolution, err error) {

    data, err := ioutil.ReadFile (filePath)
    if err != nil {
        return
    }

    lines := strings.Split(string(data), "\n")
    index := 0

    for ;strings.TrimSpace(lines[index]) == ""; index++ {
    }

    if strings.TrimSpace(lines[index]) != "Microsoft Visual Studio Solution File, Format Version 11.00" {
        err = errors.New ("Not a known VS Shell format: " + lines[index])
        return
    }
    index++;

    for ;strings.TrimSpace(lines[index]) == ""; index++ {
    }

    if strings.TrimSpace(lines[index]) != "# Atmel Studio Solution File, Format Version 11.00" {
        err = errors.New ("Not a known Atmel Studio Solution format: " + lines[index])
        return
    }
    index++;

    projectData := []map[string]string {}
    for ;strings.TrimSpace(lines[index]) != "EndProject"; index++ {

        matches := getMatchesAsMap (projectRegexId, lines[index])
        projectData = append (projectData, matches)
        if len (matches) == 0 {
            err = errors.New ("No matching solution/project found")
            return
        }
    }

    solution = new(VsSolution)
    projects := []*VsProject {}
    for _, p:= range (projectData) {
        project := &VsProject {Guid : p["project_guid"], SolutionRelativePath : p["project_path"], Solution : solution }
        projects = append (projects, project)
    }

    solution.Name = projectData[0]["name"]
    solution.Guid = projectData[0]["guid"]
    solution.Projects = projects
    solution.FilePath = filePath

    return solution, nil
}

func LoadSolution (filePath string) (*VsSolution, error) {
    return parseSolution (filePath)
}


