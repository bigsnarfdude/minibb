-- Wiki System Schema
-- Adapted from MiniBB forum structure

-- Namespaces table (was boards)
CREATE TABLE IF NOT EXISTS namespaces (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Pages table (was topics)
CREATE TABLE IF NOT EXISTS pages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    namespace_id INTEGER NOT NULL,
    slug TEXT NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    author TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    current_revision INTEGER DEFAULT 1,
    locked BOOLEAN DEFAULT 0,
    FOREIGN KEY (namespace_id) REFERENCES namespaces(id) ON DELETE CASCADE,
    UNIQUE(namespace_id, slug)
);

-- Revisions table (was posts)
CREATE TABLE IF NOT EXISTS revisions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    page_id INTEGER NOT NULL,
    revision_number INTEGER NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    author TEXT NOT NULL,
    edit_summary TEXT DEFAULT '',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (page_id) REFERENCES pages(id) ON DELETE CASCADE,
    UNIQUE(page_id, revision_number)
);

-- Page links table (for tracking internal links)
CREATE TABLE IF NOT EXISTS page_links (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    from_page_id INTEGER NOT NULL,
    to_page_id INTEGER NOT NULL,
    link_text TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (from_page_id) REFERENCES pages(id) ON DELETE CASCADE,
    FOREIGN KEY (to_page_id) REFERENCES pages(id) ON DELETE CASCADE,
    UNIQUE(from_page_id, to_page_id)
);

-- Comments table (for discussion pages)
CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    page_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    author TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (page_id) REFERENCES pages(id) ON DELETE CASCADE
);

-- Indexes for performance
CREATE INDEX IF NOT EXISTS idx_pages_namespace_id ON pages(namespace_id);
CREATE INDEX IF NOT EXISTS idx_pages_slug ON pages(slug);
CREATE INDEX IF NOT EXISTS idx_pages_updated_at ON pages(updated_at);
CREATE INDEX IF NOT EXISTS idx_revisions_page_id ON revisions(page_id);
CREATE INDEX IF NOT EXISTS idx_revisions_revision_number ON revisions(revision_number);
CREATE INDEX IF NOT EXISTS idx_page_links_from_page_id ON page_links(from_page_id);
CREATE INDEX IF NOT EXISTS idx_page_links_to_page_id ON page_links(to_page_id);
CREATE INDEX IF NOT EXISTS idx_comments_page_id ON comments(page_id);

-- Full-text search index for content
CREATE VIRTUAL TABLE IF NOT EXISTS pages_fts USING fts5(
    title, content, author, content=pages, content_rowid=id
);

-- Triggers to update timestamps and search index
CREATE TRIGGER IF NOT EXISTS update_pages_updated_at
    AFTER UPDATE ON pages
    FOR EACH ROW
    BEGIN
        UPDATE pages SET updated_at = CURRENT_TIMESTAMP WHERE id = NEW.id;
    END;

CREATE TRIGGER IF NOT EXISTS pages_fts_insert
    AFTER INSERT ON pages
    FOR EACH ROW
    BEGIN
        INSERT INTO pages_fts(rowid, title, content, author) 
        VALUES (NEW.id, NEW.title, NEW.content, NEW.author);
    END;

CREATE TRIGGER IF NOT EXISTS pages_fts_update
    AFTER UPDATE ON pages
    FOR EACH ROW
    BEGIN
        UPDATE pages_fts SET title = NEW.title, content = NEW.content, author = NEW.author 
        WHERE rowid = NEW.id;
    END;

CREATE TRIGGER IF NOT EXISTS pages_fts_delete
    AFTER DELETE ON pages
    FOR EACH ROW
    BEGIN
        DELETE FROM pages_fts WHERE rowid = OLD.id;
    END;