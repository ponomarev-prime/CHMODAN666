# CHMODAN

Однажды у меня спросили:

[How to fix "chmod" permissions after running "chmod 222 /bin/chmod"?](https://askubuntu.com/questions/1483230/how-to-fix-chmod-permissions-after-running-chmod-222-bin-chmod)

И я ответил: 

```
os.Chmod(fileName, 0666)
```

## Как пользоваться

```
> go build chmodan.go
```

```
> chmod 222 example.txt 
> stat -c %a example.txt 
222
> ./chmodan 
Права доступа к файлу успешно изменены.
> stat -c %a example.txt 
666
```

Если что-то пойдёт не так, проверь:
```
export PATH=$PATH:/usr/local/go/bin
```

## Real `chmod`

Изменим права к системному `chmod`:
```
> chmod --version
chmod (GNU coreutils) 8.32
...
> which chmod
/usr/bin/chmod
> sudo chmod 222 /usr/bin/chmod
[sudo] password for XXXXXXX: 
> chmod --version
bash: /usr/bin/chmod: Permission denied
> sudo chmod 755 /usr/bin/chmod
sudo: chmod: command not found
```
Системный `chmod` успешно сломан.

Изменим код:
```
...
	fileName := "/usr/bin/chmod"

	// Установка прав доступа 755
	err := os.Chmod(fileName, 0755)
...
```

Перекомпилируем и запустим приложение:
```
> go build chmodan
> stat -c %a /usr/bin/chmod
222
> sudo ./chmodan 
Права доступа к файлу успешно изменены.
> stat -c %a /usr/bin/chmod
755
> chmod --version
chmod (GNU coreutils) 8.32
```