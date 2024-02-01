@echo off
rem
rem Ejemplo de manejo de la fecha y hora actual - v2014-02-05
rem

chcp 1252 > NUL

setlocal

git add .
set TEXTO=%1% %DATE% - %TIME%

git commit -m "%TEXTO%"

echo DONE COMMIT: %TEXTO%
