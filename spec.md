Nanomail Specification (v0.0.0 DRAFT)
==============

1 Overview 
----------

Nanomail is a text-based internet communication protocol, inspired by email. 

A nanomail server runs on port 1999 and handles receiving messages. A nanomail
client signs messages and sends them to the destination server.

Nanomail combines many technologies that make up email infrastructure, like
SMTP, POP and DMARC into a single, dramatically simpler specification.

Nanomail is designed around personal, text-based communication, as well as for
online public communication, such as discussion boards. It should be simple
enough for a basic implementation to be easy, while thorough enough for public
and private communication.

All requests should be encased in TLS.

Structure of a message
---------------------

Nanomail messages look a lot like email messages. They have these
headers, in the following order, followed by an empty line:

```
Signature: abcdef123
From: sally@example.com
To: bob@gmail.com
Sent-At: 2022-03-01T01:23:45Z
Thread-Id: def789
Subject: My Email

Message body
```

Header values MUST NOT contain newline characters. Header names MUST include a
space after the colon. Header names are not case sensitive. Headers SHOULD NOT
be duplicated, duplicate headers may be discarded. Header order after the
signature is not relevant

The body of a nanomail message MUST consist of UTF-8 formatted gemtext, as
described in the  [Gemini specification](https://gemini.circumlunar.space/docs/specification.gmi).

Email addresses must be of the format:

```
[some-text]@[valid URI]
```

Some-text can't include an @ symbol. TBD how to define a URI

Sent-at is a datetime in UTC that must follow this RFC-3339 format:
YYYY-MM-DDTHH:MM:SSZ.

Unlike in email, sent-at should be relatively close to received-at -- each attempt at delivery should update the sent-at field and re-sign the message.

Thread-Id represents the thread that a message may be a part of. Threads
consist of a linear, not tree-based structure. 

Signatures, or some truncated form of it, may be used as a 'message ID', e.g.
for generating hyperlinks in a web forum.

Threading
----------

Clients SHOULD NOT change subjects for messages that are in the same thread.

Threads SHOULD be sorted by Sent-At.


Signing the message
------------------

Authentication is handled outside of the nanomail system. You are responsible for A. saving a message and B. putting your public key somewhere public.

This is sort of like DMARC, but uses the server instead of DNS records.

We use an asymmetric key algorithm to verify the integrity and
authenticity of a message.

Signatures should use ED25519.

The signature signs everything after the newline at the end of the signature.

Nanomail servers SHOULD verify that the Sent-At header of the message is within
some reasonable time period.

Sending mail 
------------

A nanomail message represents one-to-one communication between a single sender
and a single recepient. Think like physical mail in this anology.

The nanomail client is solely responsible for sending messages to a server, similar to SMTP.

No multiple recipients here. Think like physical mail

A client request consists of a command (in all caps) followed by CRLF and a message.

```
SEND
Signature: abcdef123
From: sally@example.com
To: bob@whatever.example
Subject: My Email

Hello bob. Thanks for reading my email.
```


whatever.example then makes a request to the nanomail server at example.com (if it doesnt have the key already):

```
GETKEY sally
```

Which fetches the public keyy and validates the signature. OR maybe it returns
a URL where you can get the key? which could be either gemini:// or https:// or
something else? this way, you can update your key without talking to the
server? 

This URL looks for the string nmail:somekey123 on the page. This is the cerca
model of auth https://github.com/cblgh/cerca

If key validation works, server will respond OK, else some error code

status CODES (not complete):
```
20 OK
40 INVALID REQUEST
41 USER DNE
51 SIGNATURE INVALID
```

etc

Receiving mail
-------------

Signature validation TBD

Servers SHOULD validate that the Sent-At time is reasonably accurate, ie,
consistent with a reasonable time that the message would take to be delivered.


Pulling mail (client-server)
---------------------------

Use private key authentication

```
FETCH sally [signature]
```

Returns messages if there are any. Each request returns a message, or DONE

Servers SHOULD delete the message after it is fetched.

Registering account (client-server)
-------------------

REGISTER uname (URL)

If URL is one of the allowlisted URL hosts and uname DNE, you are registered. 

The URL should be a place containing the public ed25519 key (e.g. a personal website)

You should keep this URL updated with your public key so that you can reliably send messages

Handling spam
-------------

No one will use this, so no worries about spam. If they do, consider a server
allowlist or blocklist (like how fediverse handles these things). Nanomail is a
"human-scale" technology: if your server becomes large enough that you cannot
adequately moderate it, you should cap signups. Encourage self-hosting and
single-user or few-user instances.

Server administrators should reserve postmaster@[host] to respond to any problems or abuse
