# These templates just are an easy way to auto fill the files every day. 

For this to work you have to add this to your init.lua (if using neo vim):

```lua  
vim.api.nvim_create_autocmd("BufNewFile", {
  pattern = "*journal.md",
  command = "0r /path/to/templates/journal.md"
})

```
