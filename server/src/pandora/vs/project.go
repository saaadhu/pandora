package vs

import (
    "encoding/xml"
    "io/ioutil"
    "path"
)

type Toolchain struct {
    GenerateHex string `xml:avrgcc.common.outputfiles.hex`
    GenerateLss string `xml:avrgcc.common.outputfiles.lss`
}

type toolchainSetting struct {
    toolchain Toolchain
}

type PropertyGroup struct {
    SchemaVersion string
    ProjectVersion string
    ToolchainName string
    avrdevice string
    avrdeviceseries string
    OutputType string
    Language string
    OutputFileName string
    OutputFileExtension string
    OutputDirectory string
    Name string
    ToolchainSetting toolchainSetting
}

type CompileItem struct {
    Name string `xml:"Include,attr"`
}

type itemgroup struct {
    CompileItems []CompileItem `xml:"Compile"`
}

type Project struct {
   PropertyGroups []PropertyGroup `xml:"PropertyGroup"`
   ItemGroups []itemgroup `xml:"ItemGroup"`
}

type VsProject struct {
    Name string
    Guid string
    SolutionRelativePath string
    Solution *VsSolution
    Items []string
}

func (p *VsProject) Load () (err error) {
    contents, err := ioutil.ReadFile (path.Dir(p.Solution.FilePath) + "/" + p.SolutionRelativePath)
    if err != nil {
        return err
    }

    pData := Project{}
    xml.Unmarshal (contents, &pData)
    p.Name = pData.PropertyGroups[0].Name

    items := []string {}
    for _,item := range (pData.ItemGroups) {
        for _, i := range (item.CompileItems) {
           items = append (items, i.Name)
       }
    }

    p.Items = items
    return
}
