-- Default namespaces and pages for Wiki system

-- Default namespaces
INSERT INTO namespaces (slug, name, description) VALUES
('main', 'Main', 'Main articles and content'),
('help', 'Help', 'Documentation and help articles'),
('templates', 'Templates', 'Reusable page templates'),
('project', 'Project', 'Project-related documentation');

-- Default pages
INSERT INTO pages (namespace_id, slug, title, content, author, current_revision) VALUES
(1, 'welcome', 'Welcome to the Wiki', 
'# Welcome to the Wiki

This is your wiki homepage. You can edit this page to customize your wiki.

## Getting Started

- [[Create a New Page|help:creating-pages]]
- [[Editing Guide|help:editing]]
- [[Markdown Syntax|help:markdown]]

## Recent Changes

Check the recent changes to see what''s been updated.

## Popular Pages

- [[Project Overview|project:overview]]
- [[Team Members|main:team]]
- [[Meeting Notes|project:meetings]]

---

*This wiki is powered by a custom wiki engine built with Go and React.*', 
'Admin', 1),

(2, 'creating-pages', 'Creating Pages',
'# Creating Pages

To create a new page, simply create a link to it using the wiki link syntax:

```
[[Page Title]]
or
[[Page Title|namespace:page-slug]]
```

## Page Naming

- Use descriptive titles
- Avoid special characters
- Use hyphens for spaces in slugs

## Page Structure

Start with a clear heading and organize content with subheadings:

```markdown
# Main Title

## Section 1

Content here...

## Section 2

More content...
```

## Linking

- Internal links: `[[Page Title]]`
- External links: `[Link Text](http://example.com)`
- Cross-namespace: `[[Page Title|namespace:slug]]`',
'Admin', 1),

(2, 'editing', 'Editing Guide',
'# Editing Guide

This wiki supports full Markdown syntax for formatting.

## Basic Formatting

- **Bold text** with `**text**`
- *Italic text* with `*text*`
- `Code` with backticks
- ~~Strikethrough~~ with `~~text~~`

## Lists

Unordered lists:
- Item 1
- Item 2
  - Nested item

Ordered lists:
1. First item
2. Second item
3. Third item

## Code Blocks

```javascript
function example() {
    return "Hello, World!";
}
```

## Tables

| Header 1 | Header 2 |
|----------|----------|
| Cell 1   | Cell 2   |
| Cell 3   | Cell 4   |

## Images

![Alt text](image-url.jpg)

## Blockquotes

> This is a blockquote.
> It can span multiple lines.

## Horizontal Rules

---

## Wiki Links

- [[Internal Page]]
- [[Display Text|actual-page-slug]]
- [[Page in Namespace|namespace:page-slug]]',
'Admin', 1),

(2, 'markdown', 'Markdown Syntax',
'# Markdown Syntax Reference

This is a quick reference for Markdown syntax supported in this wiki.

## Headers

```
# H1 Header
## H2 Header  
### H3 Header
#### H4 Header
##### H5 Header
###### H6 Header
```

## Emphasis

- *Italic* or _italic_
- **Bold** or __bold__
- ***Bold and italic***
- ~~Strikethrough~~

## Links and Images

- [Link text](http://example.com)
- ![Image alt text](image.jpg)
- [[Wiki Link]]
- [[Wiki Link|namespace:page-slug]]

## Lists

### Unordered
- Item 1
- Item 2
  - Nested item
  - Another nested item

### Ordered
1. First item
2. Second item
3. Third item

## Code

Inline `code` with backticks.

Code blocks:
```language
code here
```

## Tables

| Left | Center | Right |
|------|:------:|------:|
| L1   | C1     | R1    |
| L2   | C2     | R2    |

## Blockquotes

> Blockquote text
> Can span multiple lines

## Horizontal Rule

---

## Task Lists

- [x] Completed task
- [ ] Incomplete task
- [ ] Another task',
'Admin', 1),

(4, 'overview', 'Project Overview',
'# Project Overview

This page contains an overview of the current project.

## Objectives

- Goal 1: Create a robust wiki system
- Goal 2: Support collaborative editing
- Goal 3: Implement version control

## Timeline

| Phase | Description | Status |
|-------|-------------|--------|
| 1     | Planning    | ✅ Complete |
| 2     | Development | 🔄 In Progress |
| 3     | Testing     | ⏳ Pending |
| 4     | Deployment  | ⏳ Pending |

## Team

- [[John Doe|main:john-doe]] - Project Lead
- [[Jane Smith|main:jane-smith]] - Developer
- [[Bob Johnson|main:bob-johnson]] - Designer

## Resources

- [[Meeting Notes|project:meetings]]
- [[Technical Specifications|project:tech-specs]]
- [[User Stories|project:user-stories]]

## Links

- [Project Repository](https://github.com/example/project)
- [Issue Tracker](https://github.com/example/project/issues)
- [Documentation](https://docs.example.com)',
'Admin', 1);

-- Create initial revisions
INSERT INTO revisions (page_id, revision_number, title, content, author, edit_summary) VALUES
(1, 1, 'Welcome to the Wiki', 
'# Welcome to the Wiki

This is your wiki homepage. You can edit this page to customize your wiki.

## Getting Started

- [[Create a New Page|help:creating-pages]]
- [[Editing Guide|help:editing]]
- [[Markdown Syntax|help:markdown]]

## Recent Changes

Check the recent changes to see what''s been updated.

## Popular Pages

- [[Project Overview|project:overview]]
- [[Team Members|main:team]]
- [[Meeting Notes|project:meetings]]

---

*This wiki is powered by a custom wiki engine built with Go and React.*', 
'Admin', 'Initial page creation'),

(2, 1, 'Creating Pages',
'# Creating Pages

To create a new page, simply create a link to it using the wiki link syntax:

```
[[Page Title]]
or
[[Page Title|namespace:page-slug]]
```

## Page Naming

- Use descriptive titles
- Avoid special characters
- Use hyphens for spaces in slugs

## Page Structure

Start with a clear heading and organize content with subheadings:

```markdown
# Main Title

## Section 1

Content here...

## Section 2

More content...
```

## Linking

- Internal links: `[[Page Title]]`
- External links: `[Link Text](http://example.com)`
- Cross-namespace: `[[Page Title|namespace:slug]]`',
'Admin', 'Initial page creation'),

(3, 1, 'Editing Guide',
'# Editing Guide

This wiki supports full Markdown syntax for formatting.

## Basic Formatting

- **Bold text** with `**text**`
- *Italic text* with `*text*`
- `Code` with backticks
- ~~Strikethrough~~ with `~~text~~`

## Lists

Unordered lists:
- Item 1
- Item 2
  - Nested item

Ordered lists:
1. First item
2. Second item
3. Third item

## Code Blocks

```javascript
function example() {
    return "Hello, World!";
}
```

## Tables

| Header 1 | Header 2 |
|----------|----------|
| Cell 1   | Cell 2   |
| Cell 3   | Cell 4   |

## Images

![Alt text](image-url.jpg)

## Blockquotes

> This is a blockquote.
> It can span multiple lines.

## Horizontal Rules

---

## Wiki Links

- [[Internal Page]]
- [[Display Text|actual-page-slug]]
- [[Page in Namespace|namespace:page-slug]]',
'Admin', 'Initial page creation'),

(4, 1, 'Markdown Syntax',
'# Markdown Syntax Reference

This is a quick reference for Markdown syntax supported in this wiki.

## Headers

```
# H1 Header
## H2 Header  
### H3 Header
#### H4 Header
##### H5 Header
###### H6 Header
```

## Emphasis

- *Italic* or _italic_
- **Bold** or __bold__
- ***Bold and italic***
- ~~Strikethrough~~

## Links and Images

- [Link text](http://example.com)
- ![Image alt text](image.jpg)
- [[Wiki Link]]
- [[Wiki Link|namespace:page-slug]]

## Lists

### Unordered
- Item 1
- Item 2
  - Nested item
  - Another nested item

### Ordered
1. First item
2. Second item
3. Third item

## Code

Inline `code` with backticks.

Code blocks:
```language
code here
```

## Tables

| Left | Center | Right |
|------|:------:|------:|
| L1   | C1     | R1    |
| L2   | C2     | R2    |

## Blockquotes

> Blockquote text
> Can span multiple lines

## Horizontal Rule

---

## Task Lists

- [x] Completed task
- [ ] Incomplete task
- [ ] Another task',
'Admin', 'Initial page creation'),

(5, 1, 'Project Overview',
'# Project Overview

This page contains an overview of the current project.

## Objectives

- Goal 1: Create a robust wiki system
- Goal 2: Support collaborative editing
- Goal 3: Implement version control

## Timeline

| Phase | Description | Status |
|-------|-------------|--------|
| 1     | Planning    | ✅ Complete |
| 2     | Development | 🔄 In Progress |
| 3     | Testing     | ⏳ Pending |
| 4     | Deployment  | ⏳ Pending |

## Team

- [[John Doe|main:john-doe]] - Project Lead
- [[Jane Smith|main:jane-smith]] - Developer
- [[Bob Johnson|main:bob-johnson]] - Designer

## Resources

- [[Meeting Notes|project:meetings]]
- [[Technical Specifications|project:tech-specs]]
- [[User Stories|project:user-stories]]

## Links

- [Project Repository](https://github.com/example/project)
- [Issue Tracker](https://github.com/example/project/issues)
- [Documentation](https://docs.example.com)',
'Admin', 'Initial page creation');