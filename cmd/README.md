# Quikassociate
When installing new applications on KDE Plasma, I tend to have the issue where new applications set themselves as default association for *every* file they support.
I don't like this kind of coup d'etat, so this application exists to quickly set the association for my preferred file handler.

You probably want to use this to do the broad strokes and set your associations more fine-grained afterwards.
Requires `xdg-mime`. 

### Commands
`quikassociate associate [application]`
Sets application as handler for all defined XDG MIME types.
`quikassociate query [application]` Dumps defined XDG MIME types