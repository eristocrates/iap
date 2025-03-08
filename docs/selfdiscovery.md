# [usecases and constraints](https://svatasimara.medium.com/domain-driven-design-part-2-model-53be4e01c9e2)

things should be modular, composable, ubiquitous
things should be configurable! to the nth degree. if a proposed abstraction allows for a corner case lean towards pursual.
that said, things should demonstrate by default! demonstrate from the context of my personal requirements as an example configuration, NOT "canonical use". show what it's like to walk, not where to go 
in demonstrating my example use, explicitly include embedded and/or frictionless access to the theory citation that inspired it. for example, suggest guidelines of reward/consequence design with citation to Succeeding with Adult ADHD
things should always be searchable, both for handling semantic errors and for discovering emergent patterns. "deleted"/expired things should especially be easy to search, for restoring from accidents or after changing one's mind
duplication and near duplication should be automatically detected, and configurably notify the user
wherever possible, navigatable via keyboard
automatically generated rollback snapshots of one's system (configuration especially) should be present by default but ofc configurable. 

for the following, assume a concept in the list below is tupled with the user action framework, concept definer, concept designer. for example, the user action framework represents something like crud but expanded from the perspective of the end user. 
just going off of crud as an example, for the concept title the user actions would be
create (add) title
read (search/find by) title
update (change) title
delete (make anonymous) title

but going beyond crud as inspiration, user actions also include
copy title (for bulk pasting elsewhere)
duplicating titles, liking titles (giving it the liked metadata), as well as bulk modification to an action (think modal text editor numbers before actions. wherever possible bulk actions should be permissible)

a designer is more about "presentation policy", allowing the user to establish style at a given "scope" that can cascade to "subscopes", or manually setting exceptions to ignore the style of it's "superscope" parent. it's hard to go into detail as presentation is more dependent on implementation details, but for one example the user should at least be able to say all titles with this tag have a color, but on top of that establish all titles with some sub string are rendered in a specific color or font ignoring the .

a definer defines "constraint logic", allowing functional configuration to the user, like enforcing max title length, or automatically prepending titles with a mask based on metadata (automating johnny decimal systems). as constraints are also a concept, they too fall under the user action framework, and can be created, duplicated, assigned in bulk, etc.
this extends all the way to the the user action framework itself, as user actions are a concept. i imagine a user populating an option menu with an action that associates a task to a tag with a relation link. I really like how grocy handles products, but i'm not looking to specifically recreate that. I similarly like how vintage story's entire game is essentially a mod of its own engine, meaning all user generated content has access to the same exact tools the creators used. 

I generated a summary of more ideas to take inspiration from, but for now understand the list is simple below for brevity in getting out thoughts, but a lot applies to these concepts

## tasks
task
title 
hierarchical path
note
tags



## semantics
relations (is a, has a, parallel, )
relations between notes
relations between tags

## representation
linear

board
configurable filters
savable filters as "presets"
view flat list
view hierarchical
navigate within and throught granularity levels (zoom in/out level vs different tasks at a given level)

## attachments
upload from device
attatch link

add metadata
add colors
define color scheme
import/export/translate color schemes

add plan (what i'd like to do)
add initiative (what i'm commiting to do)
add constraints to scheduler (what the app tries to do)
add 

track completion
track time

add time frame
add life span
define semantics
semantic search

## motivation
add reward
add consequence
