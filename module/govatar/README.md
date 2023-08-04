# govatar

Govatar is a silly little program written in Golang that generates a unique avatar, complete with a unique color scheme,
based on any text string. The main aim was to create a slightly more visually interesting version of Identicons.
Although govatars do contain all of the originally provided information like Identicons do, it is much harder
to retrieve. They are only meant for use as account avatars, not as data storage.

### Usage

To use call the CreateAvatar function, passing the width and height of the
canvas, the width and height of each "block", the user name (or any string) that you want
the govatar to be based on, a salt string used to generate hash and the desired vibrance
(recommended value of four).

Block width and height should be divisors of canvas width and height. The salt
used should be equal in length to the number of blocks being rendered (half of the total blocks). For example
on a 128 x 128 canvas with 16 x 16 blocks the salt should be a string with a length
of 32. If you want to use special characters please remember to encode them properly.
