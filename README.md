# nit

## Tool description

### Todo

- [x] implement git integration
- [x] implement ollama client
- [x] implement github client
- [x] implement sqlite manager
- [ ] implement custom markdown prompt
- [x] improve markdown parsing
- [x] use ollama stream response



### Suggested flow

```sh
git checkout -b my-branch
```
```sh
git add .
```
```sh
go run . draft
```
```sh
git commit -m $(go run . draft -l -f commit)
```
```sh
set TITLE $(go run . draft -l -f pr-title) 
set BODY $(go run . draft -l -f pr-body)
gh pr create --title "$TITLE" --body-file "$BODY"
```

