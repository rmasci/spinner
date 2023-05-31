# Spinner

Simple tool that launches a program / script that may be long running. This provides feedback to the user that it's being worked on. 

```azure
spinner runSomething
```

spinner will launch the program and provide a spinner. There are different types of spinners. Run the demo --demo to see which ones you like.

TODO: need to separate flags. Right now to run a program with flags you have to put it in quotes.
```azure
spinner -t 32 -d 25ms "myapp -l 1 -b 2 -c -e"
```
Spinner would then launch "myapp -l 1 -b 2 -c -e"