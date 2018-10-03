# todo2

Playing with Buffalo, I remade the toodo example app and took some notes.  

I cracked the toodo example toodo application here: 
https://github.com/gobuffalo/toodo

Following the instructions in the README I was able to bring the application up by:  
- installing postgres
- setting up DB user/passwd that matched what the app was expecting in database.yml
- creating db via buffalo
- running migrations via buffalo
- firing up the app via `buffalo dev`

This all worked swimmingling.  

The app is a toy app.  It's just a plain jane todo list.  
The app has simplistic user management.  You can register a user.  Then login as the user.  
Make a todo item, etc.  
I did notice that it didn't have any style really (like CSS type style).  I had seen Bootstrap mentioned in the docs but wasn't seeing it here.  

Also I ran into a bug:
After you created a user and you were actually supposedly able to see the list of items, the app was barfing with a   
```
items/index.html: line 0: items/_li.html: line 1: item: unknown identifier
```

I was glad to find this exact issue was already logged against the repo.  Surely a solution would be one mouse scroll away!
https://github.com/gobuffalo/toodo/issues/8

But no :(


So I played with it.  The problem was that when showing the list of items, the 'item' variable wasn't in the context that was available when the _li.html partial was being rendered.

It was also hard to get to the problem.  By this I mean that I would modify the template that was calling the partial, and I wasn't seeing a change when reloading the webpage.  
I beat my head against it for too long and then concluded that the `app.Resource("/items", ItemsResource{})` in actions/app.go must be generating the templates already.  So that it's basically ignoring what's on disk in the templates folder.  (I'm not postive on this aspect but pushed forward on that hunch)  

I cracked the buffalo source and could see in the generators that the current buffalo CRUD templates were newer.  They basically took a different approach.  Rather than using a for loop to call partial for each <li> they just did them all in one page.  So the whole 'item' not in the context was no longer an issue.

Since I wasn't sure how to fix the old one, I made a new todo app, poaching from the original but relying on the new (v0.12.6) buffalo to provide updated CRUD templates.

It turned out this was a great way to get my feet wet with buffalo.  
And the new app has working Bootstrap.  yay!  

Here's a rough list of what I did to make the new todo app.  

And then of course this repo, is the new app.  


Rough list of work to make todo2  
```
buffalo new todo2
cd todo2
buffalo generate resource items title body
edit models/item.go - update the schema to match what was in toodo
copy the routes from toodo/actions/app.go in todo2/actions/app.go
cp toodo/actions/user* todo2/actions/
buffalo g resource users email passwordHash password passwordConfirmation
```

at this point I just started hacking on it to get it working:  
- copied the actions/auth.go from toodo
- fixed up the migrations to match columns in the original toodo
- added a link so new users could register
- added a User div to show what user your're logged in as, with a link to logout
  ( this took forever b/c my `app.Use(SetCurrentUser)` call in actions/app.go was after my items resource declaration.  B/c of this my current_user would never be set in the context.  This issue blew my mind a bit.  I had no idea the order mattered in actions/app.go)


The rest is Boilerplate to setup the app...  





Thank you for choosing Buffalo for your web development needs.

## Database Setup

It looks like you chose to set up your application using a postgres database! Fantastic!

The first thing you need to do is open up the "database.yml" file and edit it to use the correct usernames, passwords, hosts, etc... that are appropriate for your environment.

You will also need to make sure that **you** start/install the database of your choice. Buffalo **won't** install and start postgres for you.

### Create Your Databases

Ok, so you've edited the "database.yml" file and started postgres, now Buffalo can create the databases in that file for you:

	$ buffalo db create -a

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.

Good luck!

[Powered by Buffalo](http://gobuffalo.io)
