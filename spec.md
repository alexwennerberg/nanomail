Nanomail Spec
==============

A lighter-weight email replacement

Deliberitly limited in scope for personal, text-based communication

Server runs on port 1999

Strucure of a message
---------------------

nanomail messages look a lot like email, but without the 50 years of annoying junk

mail messages have a fixed set of exactly 4 headers:

Signature: abcdef123
From: sally@example.com
To: bob@gmail.com
Subject: My Email

Followed by an empty newline, then the body of the email.

Header values cannoy contain newline characters.

Email addresses must be of the format:

[some-text]@[valid URI]

some-text can't include an @ symbol

The body consists of gemtext-formatted text. See the Gemini specification.

=> https://gemini.circumlunar.space/docs/specification.gmi

Signatures are (....?) Something along the lines of HTTP signatures https://datatracker.ietf.org/doc/html/draft-cavage-http-signatures-12#section-2.1.3

The signature signs everything after the newline at the end of the signature

Sending mail 
------------

Like SMTP

No multiple recipients here. Think like physical mail

Client request:

  SEND
  Signature: abcdef123
  From: sally@example.com
  To: bob@gmail.com
  Subject: My Email

  Hello bob. Thanks for reading my email.

gmail.com then makes a request to the nanomail server at example.com (if it doesnt have the key already):

  GETKEY sally

Which fetches the public keyy and validates the signature.

If key validation works, server will respond OK, else some error code

status CODES:

20 OK
40 INVALID REQUEST
41 USER DNE
51 SIGNATURE INVALID

etc

Pulling mail (client-server)
---------------------------

Like POP

  FETCH sally [signature]

Returns messages if there are any. Each request returns a message, or DONE

Handling spam
-------------

no one will use this, so no worries about spam. If they do, consider a server
allowlist or blocklist (like how fediverse handles these things)
