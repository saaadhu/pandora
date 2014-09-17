package vs
import (
    "testing"
    "io/ioutil"
)

var solutionContents string = `

Microsoft Visual Studio Solution File, Format Version 11.00                                                                                                                                                                      
# Atmel Studio Solution File, Format Version 11.00                                                              
Project("{E66E83B9-2572-4076-B26E-6BE79FF3018A}") = "TestSolution", "ShiftPWM_Non_Blocking.cppproj", "{71DB13F7-3510-4E6E-82C6-E72BF6C4C01D}"
EndProject
Global
    GlobalSection(SolutionConfigurationPlatforms) = preSolution
        Debug|AVR = Debug|AVR
        Release|AVR = Release|AVR
    EndGlobalSection
    GlobalSection(ProjectConfigurationPlatforms) = postSolution                                                 
        {71DB13F7-3510-4E6E-82C6-E72BF6C4C01D}.Debug|AVR.ActiveCfg = Debug|AVR                                  
        {71DB13F7-3510-4E6E-82C6-E72BF6C4C01D}.Debug|AVR.Build.0 = Debug|AVR                                    
        {71DB13F7-3510-4E6E-82C6-E72BF6C4C01D}.Release|AVR.ActiveCfg = Release|AVR                              
        {71DB13F7-3510-4E6E-82C6-E72BF6C4C01D}.Release|AVR.Build.0 = Release|AVR                                
    EndGlobalSection
    GlobalSection(SolutionProperties) = preSolution
        HideSolutionNode = FALSE
    EndGlobalSection
EndGlobal
`;

func createSolution() (filePath string) {
    filePath = "/tmp/TestSolution.atsln"

    ioutil.WriteFile (filePath , []byte(solutionContents), 0777)
    ioutil.WriteFile ("/tmp/ShiftPWM_Non_Blocking.cppproj", []byte(projectContents), 0777)
    return
}

func TestSolutionLoad (t *testing.T) {
    solutionFile := createSolution()
    solution, err := LoadSolution (solutionFile)
    if err != nil {
        t.Error (err.Error())
    }

    if solution == nil {
        t.Error ("No valid solution")
    }
    if solution.Name != "TestSolution" {
        t.Error ("Solution name does not match")
    }

    if solution.Guid != "{E66E83B9-2572-4076-B26E-6BE79FF3018A}" {
        t.Error ("Solution Guid does not match : " + solution.Guid)
    }

    if len(solution.Projects) != 1 {
        t.Error ("Project count does not match")
    }

    if solution.Projects[0].Guid != "{71DB13F7-3510-4E6E-82C6-E72BF6C4C01D}" {
        t.Error ("Project guid does not match")
    }

    if solution.Projects[0].SolutionRelativePath != `ShiftPWM_Non_Blocking.cppproj` {
        t.Error ("Project path does not match")
    }
}
