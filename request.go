package clio
/**

TODO: 

req.params
This property is an array containing properties mapped to the named route "parameters". For example if you have the route /user/:name, then the "name" property is available to you as req.params.name

req.body
This property is an object containing the parsed request body. This feature is provided by the bodyParser() middleware, though other body parsing middleware may follow this convention as well. This property defaults to {} when bodyParser() is used.


req.body
This property is an object containing the parsed request body. This feature is provided by the bodyParser() middleware, though other body parsing middleware may follow this convention as well. This property defaults to {} when bodyParser() is used.

// POST user[name]=tobi&user[email]=tobi@learnboost.com
req.body.user.name
// => "tobi"

req.body.user.email
// => "tobi@learnboost.com"

// POST { "name": "tobi" }
req.body.name
// => "tobi"


req.files
This property is an object of the files uploaded. This feature is provided by the bodyParser() middleware, though other body parsing middleware may follow this convention as well. This property defaults to {} when bodyParser() is used.

For example if a file field was named "image", and a file was uploaded, req.files.image would contain the following File object:

{ size: 74643,
  path: '/tmp/8ef9c52abe857867fd0a4e9a819d1876',
  name: 'edge.png',
  type: 'image/png',
  hash: false,
  lastModifiedDate: Thu Aug 09 2012 20:07:51 GMT-0700 (PDT),
  _writeStream: 
   { path: '/tmp/8ef9c52abe857867fd0a4e9a819d1876',
     fd: 13,
     writable: false,
     flags: 'w',
     encoding: 'binary',
     mode: 438,
     bytesWritten: 74643,
     busy: false,
     _queue: [],
     _open: [Function],
     drainable: true },
  length: [Getter],
  filename: [Getter],
  mime: [Getter] }
  
The bodyParser() middleware utilizes the node-formidable module internally, and accepts the same options. An example of this is the keepExtensions formidable option, defaulting to false which in this case gives you the filename "/tmp/8ef9c52abe857867fd0a4e9a819d1876" void of the ".png" extension. To enable this, and others you may pass them to bodyParser():

app.use(express.bodyParser({ keepExtensions: true, uploadDir: '/my/files' }));

req.param(name)
Return the value of param name when present.

// ?name=tobi
req.param('name')
// => "tobi"

// POST name=tobi
req.param('name')
// => "tobi"

// /user/tobi for /user/:name 
req.param('name')
// => "tobi"
Lookup is performed in the following order:

req.params
req.body
req.query
Direct access to req.body, req.params, and req.query should be favoured for clarity - unless you truly accept input from each object.


*/