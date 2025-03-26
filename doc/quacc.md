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
- `quacc <topic>` `quacc <topic>/<subtopic>` would print out the note in std out, if the note doesn't exist it will
create a new one 
- `quacc <topic>~<some text>` by using `~` sign one could search for specific text and tags in the notes
- `quacc <topic> --edit` use the `--edit` flag to edit the note

## Implementation

### Base version 
- [x] implement the basic editing and viewing document logic 
- [ ] implement querying capability
    - [ ] implement subtopic listing
    - [ ] implement querying by keywords
- [ ] support journaling 
- [ ] ability to query notes by keywords and view them

### Future 
- [ ] text searching and printing out blocks for quick access
- [ ] tags
- [ ] backlinks


