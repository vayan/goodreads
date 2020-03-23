# Goodreads API Client [![Sanity Checks](https://github.com/vayan/goodreads/workflows/Sanity%20Checks/badge.svg)](https://github.com/vayan/goodreads/actions) [![codecov](https://codecov.io/gh/vayan/goodreads/branch/master/graph/badge.svg)](https://codecov.io/gh/vayan/goodreads)

Go wrapper to interact with Goodreads API.

Still WiP see [the progress](#Progress)


# Requirement

Go 1.14 

# Installation 

`go get github.com/vayan/goodreads`


# Getting Started

You need an API Key from Goodreads, you can get one here: https://www.goodreads.com/api/keys 


## Setup the client

``golang 
gr := goodreads.NewClient("secretapikey11")
``

## Usage example

### Search

```golang 
gr.Search(ctx, "harry potter", 1)
```



# Progress 

From https://www.goodreads.com/api

- [ ] auth.user   —   Get id of user who authorized OAuth.
- [ ] author.books   —   Paginate an author's books.
- [ ] author.show   —   Get info about an author by id.
- [ ] author_following.create   —   Follow an author.
- [ ] author_following.destroy   —   Unfollow an author.
- [ ] author_following.show   —   Show author following information.
- [ ] book.isbn_to_id   —   Get Goodreads book IDs given ISBNs.
- [ ] book.id_to_work_id   —   Get Goodreads work IDs given Goodreads book IDs.
- [ ] book.review_counts   —   Get review statistics given a list of ISBNs.
- [ ] book.show   —   Get the reviews for a book given a Goodreads book id.
- [ ] book.show_by_isbn   —   Get the reviews for a book given an ISBN.
- [ ] book.title   —   Get the reviews for a book given a title string.
- [ ] comment.create   —   Create a comment.
- [ ] comment.list   —   List comments on a subject.
- [ ] events.list   —   Events in your area.
- [ ] fanship.create   —   Become fan of an author. DEPRECATED.
- [ ] fanship.destroy   —   Stop being fan of an author. DEPRECATED.
- [ ] fanship.show   —   Show fanship information. DEPRECATED.
- [ ] followers.create   —   Follow a user.
- [ ] followers.destroy   —   Unfollow a user.
- [ ] friend.confirm_recommendation   —   Confirm or decline a friend recommendation.
- [ ] friend.confirm_request   —   Confirm or decline a friend request.
- [ ] friend.requests   —   Get friend requests.
- [ ] friends.create   —   Add a friend.
- [ ] group.join   —   Join a group.
- [ ] group.list   —   List groups for a given user.
- [ ] group.members   —   Return members of a particular group.
- [ ] group.search   —   Find a group.
- [ ] group.show   —   Get info about a group by id.
- [ ] list.book   —   Get the listopia lists for a given book.
- [ ] notifications   —   See the current user's notifications.
- [ ] owned_books.create   —   Add to books owned.
- [ ] owned_books.list   —   List books owned by a user.
- [ ] owned_books.show   —   Show an owned book.
- [ ] owned_books.update   —   Update an owned book.
- [ ] owned_books.destroy   —   Delete an owned book.
- [ ] quotes.create   —   Add a quote.
- [ ] rating.create   —   Like a resource.
- [ ] rating.destroy   —   Unlike a resource.
- [ ] read_statuses.show   —   Get a user's read status.
- [ ] recommendations.show   —   Get a recommendation from a user to another user.
- [ ] review.create   —   Add review.
- [ ] review.edit   —   Edit a review.
- [ ] review.destroy   —   Delete a book review.
- [ ] reviews.list   —   Get the books on a members shelf.
- [ ] review.recent_reviews   —   Recent reviews from all members..
- [ ] review.show   —   Get a review.
- [ ] review.show_by_user_and_book   —   Get a user's review for a given book.
- [ ] search.authors   —   Find an author by name.
- [x] search.books   —   Find books by title, author, or ISBN.
- [x] series.show   —   See a series.
- [x] series.list   —   See all series by an author.
- [x] series.work   —   See all series a work is in.
- [ ] shelves.add_to_shelf   —   Add a book to a shelf.
- [ ] shelves.add_books_to_shelves   —   Add books to many shelves.
- [ ] shelves.list   —   Get a user's shelves.
- [ ] topic.create   —   Create a new topic via OAuth.
- [ ] topic.group_folder   —   Get list of topics in a group's folder.
- [ ] topic.show   —   Get info about a topic by id.
- [ ] topic.unread_group   —   Get a list of topics with unread comments.
- [ ] updates.friends   —   Get your friend updates.
- [ ] user_shelves.create   —   Add book shelf.
- [ ] user_shelves.update   —   Edit book shelf.
- [ ] user.show   —   Get info about a member by id or username.
- [ ] user.compare   —   Compare books with another member.
- [ ] user.followers   —   Get a user's followers.
- [ ] user.following   —   Get people a user is following.
- [ ] user.friends   —   Get a user's friends.
- [ ] user_status.create   —   Update user status.
- [ ] user_status.destroy   —   Delete user status.
- [ ] user_status.show   —   Get a user status.
- [ ] user_status.index   —   View user statuses.
- [ ] work.editions   —   See all editions by work.
