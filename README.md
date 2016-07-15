# Exercism-GoLang
I started TDD in GoLang using Exercism.io today: For more detailed info about the Go track see the [Exercism.io Go Lang](http://exercism.io/languages/go).

# How to start
1. It took some time to set up the CLI and run test code, I followed the instruction from here: http://exercism.io/how-it-works/newbie
2. Then I used this command to choose Golang : `exercism fetch go`
3. Then I fetched first exercise : `exercism fetch helloworld`
4. Which is saved in "configured exercises directory".
5. Later on I used this command to change my directory to keep the codes in my GoLang workspace and github: `exercism configure --dir=~C:\...``

## Test-Driven Development (TDD)

As programmers mature, they eventually want to test their code.

Here at Exercism we simulate [Test-Driven Development](http://en.wikipedia.org/wiki/Test-driven_development) (TDD), where you write your tests before writing any functionality. The simulation comes in the form of a pre-written test suite, which will signal that you have solved the problem.

It will also provide you with a safety net to explore other solutions without breaking the functionality.

### A typical TDD workflow on Exercism:

1. Run the test file and pick one test that's failing.
2. Write some code to fix the test you picked.
3. Re-run the tests to confirm the test is now passing.
4. Repeat from step 1.
5. Submit your solution (`exercism submit /path/to/file`)

## Instructions

Submissions are encouraged to be general, within reason. Having said that, it's also important not to over-engineer a solution.

It's important to remember that the goal is to make code as expressive and readable as we can. However, solutions to the hello-world exercise will not be reviewed by a person, but by rikki- the robot, who will offer an encouraging word.

To run the tests simply run the command `go test` in the exercise directory.

If the test suite contains benchmarks, you can run these with the `-bench`
flag:

    go test -bench .

For more detailed info about the Go track see the [help
page](http://exercism.io/languages/go).

## Github commands
git init

git remote add origin https://github.com/Almazi/Exercism-GoLang.git

git add -A

git commit -m "Comment to explain the changes"

git status

git push -u origin master
