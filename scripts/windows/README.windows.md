# Start powershell as admin if necessary (optional)

```powershell
powershell -ExecutionPolicy Bypass .\start_powershell_as_admin.ps1
```

OR

```powershell
Start-Process -FilePath "powershell" -Verb RunAs -ArgumentList ("-NoExit","cd {0}" -f (Get-Location).Path)
```

# Set Execution Policy (optional)

```powershell
Set-ExecutionPolicy -ExecutionPolicy Bypass -Scope Process
```

# Run with bypass

```powershell
powershell -ExecutionPolicy Bypass <script>
```
