; Script generated by the Inno Script Studio Wizard.
; SEE THE DOCUMENTATION FOR DETAILS ON CREATING INNO SETUP SCRIPT FILES!

#define MyAppName "Keybase"
#ifndef MyAppVersion
#define MyAppVersion "1.0"
#endif
; Use semantic version to name the installer,
; but we still need the x.x.x.x version because Windows.
#ifndef MySemVersion
#define MySemVersion MyAppVersion
#endif
#define MyAppPublisher "Keybase, Inc."
#define MyAppURL "http://www.keybase.io/"
#define MyExeName "keybase.exe"
#define MyGoPath GetEnv('GOPATH')
#ifndef MyGoPath
#define MyGoPath "c:\work\"
#endif
#ifndef MyExePathName
#define MyExePathName MyGoPath + "\src\github.com\keybase\client\go\keybase\" + MyExeName
#endif
#define MyGoArch GetEnv('GOARCH')
#ifndef MyGoArch
#define MyGoArch "amd64"
#endif

[Setup]
; NOTE: The value of AppId uniquely identifies this application.
; Do not use the same AppId value in installers for other applications.
; (To generate a new GUID, click Tools | Generate GUID inside the IDE.)
AppId={{70E747DE-4E09-44B0-ACAD-784AA9D79C02}
AppName={#MyAppName}
AppVersion={#MyAppVersion}
;AppVerName={#MyAppName} {#MyAppVersion}
AppPublisher={#MyAppPublisher}
AppPublisherURL={#MyAppURL}
AppSupportURL={#MyAppURL}
AppUpdatesURL={#MyAppURL}
AppCopyright=Copyright (c) 2015, Keybase
DefaultDirName={pf}\{#MyAppName}
DefaultGroupName={#MyAppName}
AllowNoIcons=yes
OutputBaseFilename=keybase_setup_{#MySemVersion}_{#MyGoArch}
SetupIconFile={#MyGoPath}\src\github.com\keybase\keybase\public\images\favicon.ico
Compression=lzma
SolidCompression=yes
UninstallDisplayIcon={app}\keybase.exe
VersionInfoVersion={#MyAppVersion}
DisableDirPage=auto
DisableProgramGroupPage=auto
; CreateUninstallRegKey=no
; Comment this out for development
; (there doesn't seem to be a way to make it conditional)
SignTool=SignCommand


[Languages]
Name: "english"; MessagesFile: "compiler:Default.isl"

[Files]
Source: "{#MyExePathName}"; DestDir: "{app}"; Flags: ignoreversion
; NOTE: Don't use "Flags: ignoreversion" on any shared system files

[Icons]
Name: "{group}\{cm:UninstallProgram,{#MyAppName}}"; Filename: "{uninstallexe}"
Name: "{group}\{#MyAppName} CMD"; Filename: "cmd.exe"; WorkingDir: "{app}"; IconFilename: "{app}\{#MyExeName}"; Parameters: "/K ""set PATH=%PATH%;{app}"""

[Registry]
Root: "HKCU"; Subkey: "SOFTWARE\Microsoft\Windows\CurrentVersion\App Paths\Keybase.exe"; ValueType: string; ValueData: "{app}\Keybase.exe"; Flags: uninsdeletekey

[Messages]
WelcomeLabel2=This will install [name/ver] on your computer.

[Run]
Filename: "{app}\{#MyExeName}"; Parameters: "ctl watchdog"; Flags: runasoriginaluser runhidden nowait

[UninstallDelete]
Type: files; Name: "{userstartup}\{#MyAppName}.vbs"

[InstallDelete]
Type: files; Name: "{userstartup}\{#MyAppName}.vbs"

[Code]
// Simply invoking "Keybase.exe service" at startup results in an unsightly
// extra console window, so we'll emit this bit of script instead.
// (yes, this is pascal code that generates vbscript.)
// Note that we delete it at uninstall.
function CreateStartupScript(): boolean;
var
  fileName : string;
  lines : TArrayOfString;
begin
  Result := true;
  fileName := ExpandConstant('{userstartup}\{#MyAppName}.vbs');
  SetArrayLength(lines, 4);

  lines[0] := 'Dim WinScriptHost';
  lines[1] := 'Set WinScriptHost = CreateObject("WScript.Shell")';
  lines[2] := ExpandConstant('WinScriptHost.Run Chr(34) & "{app}\{#MyExeName}" & Chr(34) & " ctl watchdog", 0');
  lines[3] := 'Set WinScriptHost = Nothing';

  Result := SaveStringsToFile(filename,lines,true);
  exit;
end;

 
procedure CurStepChanged(CurStep: TSetupStep);
var
  ResultCode: Integer;

begin
  if  CurStep=ssPostInstall then
    begin
      CreateStartupScript();
    end
end;

