@echo off
rem
rem Ejemplo de manejo de la fecha y hora actual - v2014-02-05
rem

chcp 1252 > NUL

setlocal

git add .
set TEXTO=Siguiente Commit %DATE% - %TIME%

git commit -m "%TEXTO%"

