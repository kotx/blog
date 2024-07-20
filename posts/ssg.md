In the past few years, the popularity of Node.js, bundlers, and various JavaScript frameworks in web development has risen rapidly. There is a growing list of dependencies and abstractions that create distractions and friction between the developer and their creation.

So rather than dealing with all that, let's (counterintuitively) make a simple static site generator so we can focus on writing content!

## What the heck is a SSG?

Simply put, a static site generator (SSG for short) is a program that turns input files (or code) into a collection of HTML, CSS, and/or JavaScript that can be directly displayed in the browser.

### Why do I need one?

You *could* write all the HTML files yourself, but an SSG greatly enhances the maintainability of your website. For example, instead of copy-pasting common elements to make pages that look alike, you can use templates to edit them in one location! Plus, many SSGs can do bundling, image optimization, and minification for you!

## Picking the right tools

First, we need a programming language to write it in. I picked [Go](https://go.dev) due to its readability and fast prototyping, but you should choose what you are most comfortable with. Some other languages I would recommend are [Python](https://python.org), [JavaScript](https://nodejs.org), and [Ruby](https://ruby-lang.org).

Next, we'll need a templating language to craft our pages. There are lots of these, including [Jinja](https://jinja.palletsprojects.com), [Handlebars](https://handlebarsjs.com), and [ERB](https://github.com/ruby/erb). I'll be using [Templ](https://templ.guide).

## Get coding!

There's a few basic components that we will need to output HTML. We need to:

- Establish a list of files to pass to our template engine.
- Compute the variables to pass to each template, if any.
- Output the rendered templates to a directory.

### Routing

Your SSR needs operate on a list of input files. This can be as easy as storing them as keys in a map or using file-based routing.

For example, you can:
- Store it as a variable in code (the easiest method)
- Store it in a separate file (such as a `pages.json` file)
- Generate it from the filesystem structure (such as Next.js's [filesystem-based routing](https://nextjs.org/docs/pages/building-your-application/routing))

### Rendering (with metadata)

Your template engine likely has a method to pass variables to individual templates. Your SSG needs to render these templates with the appropriate metadata, and write the output to the file you specify.

As of this writing, this blog's SSG employs a strategy entirely in code:
```go
// map of output file paths to contents
routes := map[string]string{
    "index.html": pages.Index(data), // render template
    "rss.xml": func() string, error { // custom render function
        feed := &feeds.Feed{
            // ...
        }

        rss, err := feed.ToRss()
        if err != nil {
            return "", err
        }

        return rss, nil
    }(),
}
```

## Next steps:

There's still quite a lot of work to do as your needs grow, so here's a few ideas to get started:
- If your output directory contains non-generated files, have a mechanism to mark generated files. (This can be as simple as creating a generated-files.txt file)
- Generate files other than HTML pages, such as an RSS feed.
- Minify (or prettify) output.
- Re-encode images with AVIF or incorporate progressive rendering.
- Multithreaded rendering.
- Create a local development server with live-reload.
