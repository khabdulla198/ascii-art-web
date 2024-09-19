Description:
The user types an input and selects an option of the banner , generating the selected 
ascii code characters in a separate page.

Authors: 
Khadijah Abdulla 
Maryam Budawas
Sana Jamshaid

Usage: how to run 
1) go run main.go in the terminal
2) press the link thats given as a response
3) write an input in the textarea
4) select a banner style in the options provided
5) submit and receive the generated ascii art in a separate page

Implementation Details: algorithm
1) 1) we made a handler for the home page also known as the root directory (/)
2) 2) we made another handler for the result page which contains the generated ascii depending on the input
   3) we defined the cooresponding handlers - the home page will read the index.html file and print it if no error
   4) the second handler function will be responsible for generating the ascii art, and also handling the errors
   5) in the second handler, we call three functions - validate, convert and generate
   6) The three functions above are responsible for processing the input, making sure the selected option is available and converting the string to an ascii art in its respective location
