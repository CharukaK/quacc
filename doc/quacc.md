# How this thing works

## Underlying note structure
- All notes originate from the root directory and are stored as individual files.
- When a note is expanded, such as in the format `<topic>/<subtopic>`, a directory named after the topic is created. If
a root note previously existed, it will be relocated into this newly created directory.
```
root/
│
├── note1.txt
├── note2.txt
└── Topic/
    ├── subtopic1.txt
    └── subtopic2.txt
```

## Command structure

- `quacc` is the root command this will list out the root topics and a list of tags available 
- `quacc <topic>` `quacc <topic>/<subtopic>` would print out the note in std out
- `quacc <topic>~<some text>` by using `~` sign one could search for specific text and tags in the notes
- `quacc edit <topic>/<subtopic>?` to create and edit notes

## Implementation

### Base version 
- [ ] implement the basic editing and viewing document logic 
- [ ] Configuration system using a file some sort of `.conf` 

### Future 
- [ ] text searching and printing out blocks for quick access
- [ ] tags
- [ ] backlinks
- [ ] Journal/track tasks etc

