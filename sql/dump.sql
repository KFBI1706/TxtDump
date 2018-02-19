--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.11
-- Dumped by pg_dump version 9.5.11

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: text; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE text (
    id integer NOT NULL,
    pubid integer NOT NULL,
    text pg_catalog.text NOT NULL,
    created_at timestamp without time zone,
    title character varying,
    editid integer
);


ALTER TABLE text OWNER TO postgres;

--
-- Name: text_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE text_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE text_id_seq OWNER TO postgres;

--
-- Name: text_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE text_id_seq OWNED BY text.id;


--
-- Name: id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY text ALTER COLUMN id SET DEFAULT nextval('text_id_seq'::regclass);


--
-- Data for Name: text; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY text (id, pubid, text, created_at, title, editid) FROM stdin;
2	8277713	An h1 header\r\n============\r\n\r\nParagraphs are separated by a blank line.\r\n\r\n2nd paragraph. *Italic*, **bold**, and `monospace`. Itemized lists\r\nlook like:\r\n\r\n  * this one\r\n  * that one\r\n  * the other one\r\n\r\nNote that --- not considering the asterisk --- the actual text\r\ncontent starts at 4-columns in.\r\n\r\n> Block quotes are\r\n> written like so.\r\n>\r\n> They can span multiple paragraphs,\r\n> if you like.\r\n\r\nUse 3 dashes for an em-dash. Use 2 dashes for ranges (ex., "it's all\r\nin chapters 12--14"). Three dots ... will be converted to an ellipsis.\r\nUnicode is supported. ☺\r\n\r\n\r\n\r\nAn h2 header\r\n------------\r\n\r\nHere's a numbered list:\r\n\r\n 1. first item\r\n 2. second item\r\n 3. third item\r\n\r\nNote again how the actual text starts at 4 columns in (4 characters\r\nfrom the left side). Here's a code sample:\r\n\r\n    # Let me re-iterate ...\r\n    for i in 1 .. 10 { do-something(i) }\r\n\r\nAs you probably guessed, indented 4 spaces. By the way, instead of\r\nindenting the block, you can use delimited blocks, if you like:\r\n\r\n~~~\r\ndefine foobar() {\r\n    print "Welcome to flavor country!";\r\n}\r\n~~~\r\n\r\n(which makes copying & pasting easier). You can optionally mark the\r\ndelimited block for Pandoc to syntax highlight it:\r\n\r\n~~~python\r\nimport time\r\n# Quick, count to ten!\r\nfor i in range(10):\r\n    # (but not *too* quick)\r\n    time.sleep(0.5)\r\n    print i\r\n~~~\r\n\r\n\r\n\r\n### An h3 header ###\r\n\r\nNow a nested list:\r\n\r\n 1. First, get these ingredients:\r\n\r\n      * carrots\r\n      * celery\r\n      * lentils\r\n\r\n 2. Boil some water.\r\n\r\n 3. Dump everything in the pot and follow\r\n    this algorithm:\r\n\r\n        find wooden spoon\r\n        uncover pot\r\n        stir\r\n        cover pot\r\n        balance wooden spoon precariously on pot handle\r\n        wait 10 minutes\r\n        goto first step (or shut off burner when done)\r\n\r\n    Do not bump wooden spoon or it will fall.\r\n\r\nNotice again how text always lines up on 4-space indents (including\r\nthat last line which continues item 3 above).\r\n\r\nHere's a link to [a website](http://foo.bar), to a [local\r\ndoc](local-doc.html), and to a [section heading in the current\r\ndoc](#an-h2-header). Here's a footnote [^1].\r\n\r\n[^1]: Footnote text goes here.\r\n\r\nTables can look like this:\r\n\r\nsize  material      color\r\n----  ------------  ------------\r\n9     leather       brown\r\n10    hemp canvas   natural\r\n11    glass         transparent\r\n\r\nTable: Shoes, their sizes, and what they're made of\r\n\r\n(The above is the caption for the table.) Pandoc also supports\r\nmulti-line tables:\r\n\r\n--------  -----------------------\r\nkeyword   text\r\n--------  -----------------------\r\nred       Sunsets, apples, and\r\n          other red or reddish\r\n          things.\r\n\r\ngreen     Leaves, grass, frogs\r\n          and other things it's\r\n          not easy being.\r\n--------  -----------------------\r\n\r\nA horizontal rule follows.\r\n\r\n***\r\n\r\nHere's a definition list:\r\n\r\napples\r\n  : Good for making applesauce.\r\noranges\r\n  : Citrus!\r\ntomatoes\r\n  : There's no "e" in tomatoe.\r\n\r\nAgain, text is indented 4 spaces. (Put a blank line between each\r\nterm/definition pair to spread things out more.)\r\n\r\nHere's a "line block":\r\n\r\n| Line one\r\n|   Line too\r\n| Line tree\r\n\r\nand images can be specified like so:\r\n\r\n![example image](example-image.jpg "An exemplary image")\r\n\r\nInline math equations go in like so: $\\omega = d\\phi / dt$. Display\r\nmath should get its own line and be put in in double-dollarsigns:\r\n\r\n$$I = \\int \\rho R^{2} dV$$\r\n\r\nAnd note that you can backslash-escape any punctuation characters	2018-01-09 00:00:00		6659175
8	1317501	# TxtDump\n\n\n### Routes (Api):\n\n#### Request a post:\n```\nHOST/api/v1/post/{id}/request | Returns the post id and content \n```\nExample:\n```json\ncurl http://localhost:1337/api/v1/post/{id}/request\n{"ID":11,"Content":"not implemented yet"}\n```\n#### Create Post with requested ID:\n\n```\nHOST/api/v1/post/create | Creates the post with the submitted ID:\n```\nExample:\n```json\ncurl -H "Content-Type: application/json" -X POST -d '{"Content":"I really hate javascript"}' http://localhost:1337/api/v1/post/create \nReturns: {"PubID":9175728,"Content":"I really hate Javascript","Sucsess":true,"Time":""}\n```\n\nDbstring example:\nThe program looks for a file named dbstring when running this is then converted into the info used to connect to the DB for more info about this read: https://godoc.org/github.com/lib/pq\n```\nuser=postgres dbname=db password=12345 host=HOSTIP\n```	2018-02-06 12:16:48.052239	XD	6659175
7	9574364	Ｓ　Ａ　Ｎ　Ｄ　Ｅ　Ｒ　Ｓ　Ｉ　Ｖ　Ｅ　Ｒ　Ｔ　Ｓ　Ｅ　Ｎ　サニタひ翁め誕違隠気猿ュ価 ピ	2018-02-05 13:58:42.355939	P-POST KASSe	6659175
6	1809685	Det er ikke mass igjen	2018-02-05 12:56:54.680673	test	6659175
5	7159497	XXD\n	2018-02-05 12:01:06.129466	Anusarbeidet	6659175
4	2834906	lmao	2018-02-02 00:00:00	xd	6659175
12	5580586	text	2018-02-13 15:53:05.744332	Title	4553279
3	4945485	Short	2018-01-10 00:00:00	\N	6659175
13	8585425	# TxtDump\r\n\r\n## All current routes:\r\n```\r\n/\r\n/api/v1/post/amount\r\n/api/v1/post/{id}/request\r\n/api/v1/post/create\r\n/api/v1/post/{id}/edit/{editid}\r\n/api/v1/post/{id}/delete/{editid}\r\n/post/{id}/request\r\n/post/{id}/edit/{editid}\r\n/post/{id}/edit/{editid}/post\r\n/post/{id}/delete/{editid}/\r\n/post/create\r\n/post/create/new\r\n/static/\r\n```\r\n\r\n## Routes (Api):\r\n\r\n#### Request a post:\r\n```\r\nHOST/api/v1/post/{id}/request | Returns the post id and content \r\n```\r\nExample:\r\n```json\r\ncurl http://localhost:1337/api/v1/post/{id}/request\r\n```\r\nReturns:\r\n```json\r\n{"PubID":9175728,"EditID":0,"Content":"I really REALLY hate javascript","Title":"Dette e ein title","Sucsess":true,"Time":"2017-12-27T00:00:00Z"}\r\n```\r\n#### Create Post:\r\n\r\n```\r\nHOST/api/v1/post/create | Creates the post:\r\n```\r\nExample:\r\n```json\r\ncurl -H "Content-Type: application/json" -X POST -d '{"Title":"Title","Content":"text"}' http://localhost:1337/api/v1/post/create\r\n```\r\nResponse:\r\n```json\r\n{"PubID":5580586,"EditID":4553279,"Content":"text","Title":"Title","Sucsess":true,"Time":"0001-01-01T00:00:00Z"}\r\n```\r\n#### Edit Post:\r\n\r\n```\r\nHOST/api/v1/post/{pubid}/edit/{editid} | Edits the post:\r\n```\r\n```json\r\ncurl -H "Content-Type: application/json" -X POST -d '{"Title":"lmao",\r\n"Content":"tyest"}' http://localhost:1337/api/v1/post/4750794/edit/8986640\r\n```\r\n#### Delete Post:\r\n\r\n```\r\nHOST/api/v1/post/{pubid}/delete/{editid} | Deletes the post:\r\n```\r\nExample\r\n```\r\ncurl http://localhost:1337/api/v1/post/4750794/delete/8986640\r\n```\r\n<br>\r\n\r\n#### Dbstring example:\r\nThe program looks for a file named dbstring when running this is then converted into the info used to connect to the DB for more info about this read: https://godoc.org/github.com/lib/pq\r\n```\r\nuser=postgres dbname=db password=12345 host=HOSTIP\r\n```	2018-02-16 11:24:34.374044	Why React developers should modularize their applications?	6437975
1	9175728	# TxtDump\n## All current routes:\n```\n/\n/api/v1/post/amount\n/api/v1/post/{id}/request\n/api/v1/post/create\n/api/v1/post/{id}/edit/{editid}\n/api/v1/post/{id}/delete/{editid}\n/post/{id}/request\n/post/{id}/edit/{editid}\n/post/{id}/edit/{editid}/post\n/post/{id}/delete/{editid}/\n/post/create\n/post/create/new\n/static/\n```\n## Routes (Api):\n#### Request a post:\n```\nHOST/api/v1/post/{id}/request | Returns the post id and content \n```\nExample:\n```json\ncurl http://localhost:1337/api/v1/post/{id}/request\n```\nReturns:\n```json\n{"PubID":9175728,"EditID":0,"Content":"I really REALLY hate javascript","Title":"Dette e ein title","Sucsess":true,"Time":"2017-12-27T00:00:00Z"}\n```\n#### Create Post:\n```\nHOST/api/v1/post/create | Creates the post:\n```\nExample:\n```json\ncurl -H "Content-Type: application/json" -X POST -d '{"Title":"Title","Content":"text"}' http://localhost:1337/api/v1/post/create\n```\nResponse:\n```json\n{"PubID":5580586,"EditID":4553279,"Content":"text","Title":"Title","Sucsess":true,"Time":"0001-01-01T00:00:00Z"}\n```\n#### Edit Post:\n```\nHOST/api/v1/post/{pubid}/edit/{editid} | Edits the post:\n```\n```json\ncurl -H "Content-Type: application/json" -X POST -d '{"Title":"lmao",\n"Content":"tyest"}' http://localhost:1337/api/v1/post/4750794/edit/8986640\n```\n#### Delete Post:\n```\nHOST/api/v1/post/{pubid}/delete/{editid} | Deletes the post:\n```\nExample:\n```bash\ncurl http://localhost:1337/api/v1/post/4750794/delete/8986640\n```\n#### Dbstring example:\nThe program looks for a file named dbstring when running this is then converted into the info used to connect to the DB for more info about this read: https://godoc.org/github.com/lib/pq\n```\nuser=postgres dbname=db password=12345 host=HOSTIP\n```\n	2017-12-27 00:00:00		6659175
14	9615548	# Contributing to Go\n\nGo is an open source project.\n\nIt is the work of hundreds of contributors. We appreciate your help!\n\n## Before filing an issue\n\nIf you are unsure whether you have found a bug, please consider asking in the [golang-nuts mailing\nlist](https://groups.google.com/forum/#!forum/golang-nuts) or [other forums](https://golang.org/help/) first. If\nthe behavior you are seeing is confirmed as a bug or issue, it can easily be re-raised in the issue tracker.\n\n## Filing issues\n\nSensitive security-related issues should be reported to [security@golang.org](mailto:security@golang.org).\nSee the [security policy](https://golang.org/security) for details.\n\nThe recommended way to file an issue is by running `go bug`.\nOtherwise, when filing an issue, make sure to answer these five questions:\n\n1. What version of Go are you using (`go version`)?\n2. What operating system and processor architecture are you using?\n3. What did you do?\n4. What did you expect to see?\n5. What did you see instead?\n\nFor change proposals, see [Proposing Changes To Go](https://github.com/golang/proposal/).\n\n## Contributing code\n\nPlease read the [Contribution Guidelines](https://golang.org/doc/contribute.html) before sending patches.\n\nUnless otherwise noted, the Go source files are distributed under\nthe BSD-style license found in the LICENSE file.\n	2018-02-16 15:54:36.809221	9072415
\.


--
-- Name: text_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('text_id_seq', 14, true);


--
-- Name: text_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY text
    ADD CONSTRAINT text_pkey PRIMARY KEY (id);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

