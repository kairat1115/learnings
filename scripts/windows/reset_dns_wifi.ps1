#
# REQUIRES ADMIN PRIVILEGES
#

$wifi_adapter_mac = "38-FC-98-0F-B8-7E"

$adapter = Get-NetAdapter -IncludeHidden -Physical | Where-Object -Property MacAddress -eq $wifi_adapter_mac
# echo $adapter.Name

# Set-DnsClientServerAddress -InterfaceAlias $adapter.Name -ServerAddresses ("1.1.1.1")
Set-DnsClientServerAddress -InterfaceAlias $adapter.Name -ResetServerAddresses

# Set-DnsClientServerAddress -InterfaceAlias Wi-Fi -ServerAddresses ("1.1.1.1")
# Set-DnsClientServerAddress -InterfaceAlias Wi-Fi -ResetServerAddresses

# $pressanykey = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
