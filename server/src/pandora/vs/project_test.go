package vs

import (
    "testing"
)

func TestProjectLoad (t *testing.T) {
    solutionFile := createSolution()
    solution, err := LoadSolution (solutionFile)
    if err != nil {
        t.Error (err.Error())
    }

    project := solution.Projects[0]
    err = project.Load()

    if err != nil {
        t.Error (err.Error())
    }

    if project.Name != "ShiftPWM_Non_Blocking" {
        t.Error ("Project name does not match")
    }

    if len(project.Items) != 3 {
        t.Error ("Project item count does not match")
    }

    if project.Items[0] != "ArduinoFunctions.cpp" ||
        project.Items[1] != "main.cpp" ||
        project.Items[2] != "ShiftPWM_Non_Blocking.cpp" {
            t.Error ("Project item name does not match")
        }
}


var projectContents string = `
<?xml version="1.0" encoding="utf-8"?>
<Project DefaultTargets="Build" xmlns="http://schemas.microsoft.com/developer/msbuild/2003">
  <PropertyGroup>
    <SchemaVersion>2.0</SchemaVersion>
    <ProjectVersion>6.0</ProjectVersion>
    <ToolchainName>com.Atmel.AVRGCC8</ToolchainName>
    <ProjectGuid>{71db13f7-3510-4e6e-82c6-e72bf6c4c01d}</ProjectGuid>
    <avrdevice>ATmega32U4</avrdevice>
    <avrdeviceseries>none</avrdeviceseries>
    <OutputType>Executable</OutputType>
    <Language>CPP</Language>
    <OutputFileName>$(MSBuildProjectName)</OutputFileName>
    <OutputFileExtension>.elf</OutputFileExtension>
    <OutputDirectory>$(MSBuildProjectDirectory)\$(Configuration)</OutputDirectory>
    <AssemblyName>ShiftPWM_Non_Blocking</AssemblyName>
    <Name>ShiftPWM_Non_Blocking</Name>
    <RootNamespace>ShiftPWM_Non_Blocking</RootNamespace>
    <ToolchainFlavour>Native</ToolchainFlavour>
    <KeepTimersRunning>true</KeepTimersRunning>
    <OverrideVtor>false</OverrideVtor>
    <OverrideVtorValue />
    <eraseonlaunchrule>0</eraseonlaunchrule>
    <AsfVersion>3.1.3</AsfVersion>
  </PropertyGroup>
  <PropertyGroup Condition=" '$(Configuration)' == 'Release' ">
    <ToolchainSettings>
      <AvrGccCpp>
  <avrgcc.common.outputfiles.hex>True</avrgcc.common.outputfiles.hex>
  <avrgcc.common.outputfiles.lss>True</avrgcc.common.outputfiles.lss>
  <avrgcc.common.outputfiles.eep>True</avrgcc.common.outputfiles.eep>
  <avrgcc.compiler.general.ChangeDefaultCharTypeUnsigned>True</avrgcc.compiler.general.ChangeDefaultCharTypeUnsigned>
  <avrgcc.compiler.general.ChangeDefaultBitFieldUnsigned>True</avrgcc.compiler.general.ChangeDefaultBitFieldUnsigned>
  <avrgcc.compiler.optimization.level>Optimize for size (-Os)</avrgcc.compiler.optimization.level>
  <avrgcc.compiler.optimization.PackStructureMembers>True</avrgcc.compiler.optimization.PackStructureMembers>
  <avrgcc.compiler.optimization.AllocateBytesNeededForEnum>True</avrgcc.compiler.optimization.AllocateBytesNeededForEnum>
  <avrgcc.compiler.warnings.AllWarnings>True</avrgcc.compiler.warnings.AllWarnings>
  <avrgcccpp.compiler.general.ChangeDefaultCharTypeUnsigned>True</avrgcccpp.compiler.general.ChangeDefaultCharTypeUnsigned>
  <avrgcccpp.compiler.general.ChangeDefaultBitFieldUnsigned>True</avrgcccpp.compiler.general.ChangeDefaultBitFieldUnsigned>
  <avrgcccpp.compiler.optimization.level>Optimize for size (-Os)</avrgcccpp.compiler.optimization.level>
  <avrgcccpp.compiler.optimization.PackStructureMembers>True</avrgcccpp.compiler.optimization.PackStructureMembers>
  <avrgcccpp.compiler.optimization.AllocateBytesNeededForEnum>True</avrgcccpp.compiler.optimization.AllocateBytesNeededForEnum>
  <avrgcccpp.compiler.warnings.AllWarnings>True</avrgcccpp.compiler.warnings.AllWarnings>
  <avrgcccpp.linker.libraries.Libraries>
    <ListValues>
      <Value>m</Value>
    </ListValues>
  </avrgcccpp.linker.libraries.Libraries>
</AvrGccCpp>
    </ToolchainSettings>
  </PropertyGroup>
  <PropertyGroup Condition=" '$(Configuration)' == 'Debug' ">
    <ToolchainSettings>
      <AvrGccCpp>
  <avrgcc.common.outputfiles.hex>True</avrgcc.common.outputfiles.hex>
  <avrgcc.common.outputfiles.lss>True</avrgcc.common.outputfiles.lss>
  <avrgcc.common.outputfiles.eep>True</avrgcc.common.outputfiles.eep>
  <avrgcc.compiler.general.ChangeDefaultCharTypeUnsigned>True</avrgcc.compiler.general.ChangeDefaultCharTypeUnsigned>
  <avrgcc.compiler.general.ChangeDefaultBitFieldUnsigned>True</avrgcc.compiler.general.ChangeDefaultBitFieldUnsigned>
  <avrgcc.compiler.optimization.level>Optimize (-O1)</avrgcc.compiler.optimization.level>
  <avrgcc.compiler.optimization.PackStructureMembers>True</avrgcc.compiler.optimization.PackStructureMembers>
  <avrgcc.compiler.optimization.AllocateBytesNeededForEnum>True</avrgcc.compiler.optimization.AllocateBytesNeededForEnum>
  <avrgcc.compiler.optimization.DebugLevel>Default (-g2)</avrgcc.compiler.optimization.DebugLevel>
  <avrgcc.compiler.warnings.AllWarnings>True</avrgcc.compiler.warnings.AllWarnings>
  <avrgcccpp.compiler.general.ChangeDefaultCharTypeUnsigned>True</avrgcccpp.compiler.general.ChangeDefaultCharTypeUnsigned>
  <avrgcccpp.compiler.general.ChangeDefaultBitFieldUnsigned>True</avrgcccpp.compiler.general.ChangeDefaultBitFieldUnsigned>
  <avrgcccpp.compiler.symbols.DefSymbols>
    <ListValues>
      <Value>ARDUINO=100</Value>
      <Value>F_CPU=16000000L</Value>
    </ListValues>
  </avrgcccpp.compiler.symbols.DefSymbols>
  <avrgcccpp.compiler.directories.IncludePaths>
    <ListValues>
      <Value>D:\arduino-1.0.1\hardware\arduino\cores\arduino</Value>
      <Value>D:\arduino-1.0.1\hardware\arduino\variants\leonardo</Value>
      <Value>D:\arduino-1.0.1\libraries\ShiftPWM</Value>
    </ListValues>
  </avrgcccpp.compiler.directories.IncludePaths>
  <avrgcccpp.compiler.optimization.level>Optimize for size (-Os)</avrgcccpp.compiler.optimization.level>
  <avrgcccpp.compiler.optimization.PrepareFunctionsForGarbageCollection>True</avrgcccpp.compiler.optimization.PrepareFunctionsForGarbageCollection>
  <avrgcccpp.compiler.optimization.PackStructureMembers>True</avrgcccpp.compiler.optimization.PackStructureMembers>
  <avrgcccpp.compiler.optimization.AllocateBytesNeededForEnum>True</avrgcccpp.compiler.optimization.AllocateBytesNeededForEnum>
  <avrgcccpp.compiler.optimization.DebugLevel>Default (-g2)</avrgcccpp.compiler.optimization.DebugLevel>
  <avrgcccpp.compiler.warnings.AllWarnings>True</avrgcccpp.compiler.warnings.AllWarnings>
  <avrgcccpp.linker.libraries.Libraries>
    <ListValues>
      <Value>m</Value>
    </ListValues>
  </avrgcccpp.linker.libraries.Libraries>
  <avrgcccpp.linker.optimization.GarbageCollectUnusedSections>True</avrgcccpp.linker.optimization.GarbageCollectUnusedSections>
  <avrgcccpp.assembler.debugging.DebugLevel>Default (-Wa,-g)</avrgcccpp.assembler.debugging.DebugLevel>
</AvrGccCpp>
    </ToolchainSettings>
  </PropertyGroup>
  <ItemGroup>
    <Compile Include="ArduinoFunctions.cpp">
      <SubType>compile</SubType>
    </Compile>
    <Compile Include="main.cpp">
      <SubType>compile</SubType>
    </Compile>
    <Compile Include="ShiftPWM_Non_Blocking.cpp">
      <SubType>compile</SubType>
    </Compile>
  </ItemGroup>
  <Import Project="$(AVRSTUDIO_EXE_PATH)\\Vs\\Compiler.targets" />
</Project>
`

