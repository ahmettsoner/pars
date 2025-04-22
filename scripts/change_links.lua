function Link(el)
  -- Change markdown links to HTML links
  el.target = el.target:gsub("%.md", ".html")
  return el
end