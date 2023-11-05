package avl

var Tree *TreeNode
//TODO: Add events listeners, so the user can define themselfs what to do with the events
// Example: Maybe DebugListener to pull events and dump into console, or HttpEventListener to 
// update the web page
var TreeEvents *Queue[string]
