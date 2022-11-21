Nanomail Specification
==============

1 Overview 
----------

Nanomail is a text-based internet communication protocol, inspired by email. 

A nanomail server runs on port 1999 and handles receiving messages. A nanomail
client signs messages and sends them to the destination server, as well as 

Nanomail combines many technologies that make up email infrastructure, like
SMTP, POP and DMARC into a single, dramatically simpler specification.

Nanomail is designed around personal, text-based communication, as well as for
online public communication, such as discussion boards. It should be simple
enough for a basic implementation to be easy, while thorough enough for public
and private communication

All requests should be encased in TLS

No fancy DNS stuff. If you want a custom domain, you gotta host your own server

Only POP-like message receiving. nothing IMAP-like

Structure of a message
---------------------

Nanomail messages look a lot like email messages. They have these
headers, in the following order, followed by an empty line:

```
Signature: abcdef123
From: sally@example.com
To: bob@gmail.com
Sent-At: 2022-03-01T01:23:45
Thread-Id: def789
Subject: My Email

Message body
```

Header values MUST NOT contain newline characters. Header names MUST include a
space after the colon. Header names are not case sensitive. Headers SHOULD NOT
be dupulicated, duplicate headers may be discarded. Header order after the
signature is not relevant

The body of a nanomail message MUST consist of UTF-8 formatted gemtext, as
described in the  [Gemini specification](https://gemini.circumlunar.space/docs/specification.gmi).

Email addresses must be of the format:

```
[some-text]@[valid URI]
```

Some-text can't include an @ symbol. TBD how to define a URI

Sent-at is a RFC3339? formatted datetime in UTC. Unlike in email, sent-at should be relatively close to received-at -- each attempt at delivery should update the sent-at field and re-sign the message.

Thread-Id represents the thread that a message may be a part of. Threads
consist of a linear, not tree-based structure.

Signatures, or some truncated form of it, may be used as a 'message ID', e.g.
for generating hyperlinks in a web forum.

Signing the message
------------------

Authentication is handled outside of the nanomail system. You are responsible for A. saving a message and B. putting your public key somewhere public.

This is sort of like DMARC, but uses the server instead of DNS records.

We use an asymmetric key algorithm to verify the integrity and
authenticity of a message.

Signatures should use ED25519 (TODO Is it dumb to specify a key algo in the spec? maybe?)

The signature signs everything after the newline at the end of the signature.

Nanomail servers SHOULD verify that the Sent-At header of the message is within
some reasonable time period.

Sending mail 
------------

A nanomail message represents one-to-one communication between a single sender
and a single recepient. Think like physical mail in this anology.

The nanomail client is solely responsible for sending messages to a server, similar to SMTP.

No multiple recipients here. Think like physical mail

Client request:

```
SEND
Signature: abcdef123
From: sally@example.com
To: bob@gmail.com
Subject: My Email

Hello bob. Thanks for reading my email.
```

gmail.com then makes a request to the nanomail server at example.com (if it doesnt have the key already):

```
GETKEY sally
```

Which fetches the public keyy and validates the signature.

If key validation works, server will respond OK, else some error code

status CODES:
```
20 OK
40 INVALID REQUEST
41 USER DNE
51 SIGNATURE INVALID
```

etc

Pulling mail (client-server)
---------------------------

This is the part I'm the least sure on...

How should auth work? TLS client certs? UNAME/pw? Should this even be part of the spec? Just rsync an sqlite file or something? Leaning towards this

```
FETCH sally [signature]
```

Returns messages if there are any. Each request returns a message, or DONE

Handling spam
-------------

no one will use this, so no worries about spam. If they do, consider a server
allowlist or blocklist (like how fediverse handles these things)
