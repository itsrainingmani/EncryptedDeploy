$usr = $args[0]
$passwd = $args[1]

$credentials = New-Object System.Management.Automation.PSCredential -ArgumentList @($usr,(ConvertTo-SecureString -String $passwd -AsPlainText -Force))

Start-Process powershell -Credential $credentials -ArgumentList '-noprofile -command &{Start-Process D:\wintest.exe -verb runas}'