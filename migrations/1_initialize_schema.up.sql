CREATE TABLE IF NOT EXISTS "characters" (
	"id_char"	INTEGER NOT NULL UNIQUE,
	"name"	TEXT,
	"char_pic"	TEXT,
	PRIMARY KEY("id_char" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "llm_families" (
	"id_family"	INTEGER NOT NULL UNIQUE,
	"name"	TEXT,
	PRIMARY KEY("id_family" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "presets" (
	"id_preset"	INTEGER NOT NULL UNIQUE,
	"name"	TEXT,
	PRIMARY KEY("id_preset" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "preset_params" (
	"id_param"	INTEGER NOT NULL UNIQUE,
	"id_preset"	INTEGER NOT NULL,
	"json"	TEXT,
	PRIMARY KEY("id_param" AUTOINCREMENT),
	FOREIGN KEY("id_preset") REFERENCES "presets"("id_preset")
);
CREATE TABLE IF NOT EXISTS "chats_roles" (
	"id_role"	INTEGER NOT NULL UNIQUE,
	"name"	TEXT,
	PRIMARY KEY("id_role" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "chat_messages" (
	"id_message"	INTEGER NOT NULL UNIQUE,
	"message"	TEXT,
	"created"	INTEGER,
	"id_role"	INTEGER NOT NULL,
	"id_chat"	INTEGER NOT NULL,
	"parent"	INTEGER,
	PRIMARY KEY("id_message" AUTOINCREMENT),
	FOREIGN KEY("parent") REFERENCES "chat_messages"("id_message"),
	FOREIGN KEY("id_chat") REFERENCES "chats"("id_chat"),
	FOREIGN KEY("id_role") REFERENCES "chats_roles"("id_role")
);
CREATE TABLE IF NOT EXISTS "user_description" (
	"id_description"	INTEGER NOT NULL UNIQUE,
	"description"	TEXT,
	"tokens"	INTEGER,
	"username"	TEXT,
	"preset_name"	TEXT,
	PRIMARY KEY("id_description" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "llm_family_models" (
	"id_model"	INTEGER NOT NULL UNIQUE,
	"name"	TEXT,
	"id_family"	INTEGER NOT NULL,
	PRIMARY KEY("id_model" AUTOINCREMENT),
	FOREIGN KEY("id_family") REFERENCES "llm_families"("id_family")
);
CREATE TABLE IF NOT EXISTS "presets_prompts" (
	"id_prompt"	INTEGER NOT NULL UNIQUE,
	"json"	TEXT,
	PRIMARY KEY("id_prompt" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "chats" (
	"id_chat"	INTEGER NOT NULL UNIQUE,
	"id_char"	INTEGER NOT NULL,
	"created"	INTEGER,
	"tokens"	INTEGER,
	PRIMARY KEY("id_chat" AUTOINCREMENT),
	FOREIGN KEY("id_char") REFERENCES "characters"("id_char")
);
INSERT INTO "llm_families" VALUES (1,'openai');
INSERT INTO "llm_families" VALUES (2,'palm');
INSERT INTO "llm_families" VALUES (3,'anthropic');
INSERT INTO "chats_roles" VALUES (1,'assistant');
INSERT INTO "chats_roles" VALUES (2,'user');
INSERT INTO "chats_roles" VALUES (3,'system');
INSERT INTO "llm_family_models" VALUES (1,'gpt-4',1);
INSERT INTO "llm_family_models" VALUES (2,'gpt-3.5-turbo',1);
INSERT INTO "llm_family_models" VALUES (3,'gpt-3.5-turbo-16k',1);
INSERT INTO "llm_family_models" VALUES (4,'gpt-3.5-turbo-instruct',1);
INSERT INTO "llm_family_models" VALUES (5,'gpt-3.5-turbo-0613',1);
INSERT INTO "llm_family_models" VALUES (6,'gpt-3.5-turbo-16k-0613',1);
INSERT INTO "llm_family_models" VALUES (7,'gpt-3.5-turbo-0301',1);
INSERT INTO "llm_family_models" VALUES (8,'text-bison',2);
INSERT INTO "llm_family_models" VALUES (9,'claude-v2',3);
INSERT INTO "llm_family_models" VALUES (10,'claude-v1.3',3);
INSERT INTO "llm_family_models" VALUES (11,'claude-v1.2',3);
INSERT INTO "llm_family_models" VALUES (12,'claude-v1',3);
INSERT INTO "llm_family_models" VALUES (13,'claude-instant-100k',3);
INSERT INTO "llm_family_models" VALUES (14,'claude-instant-1.2',3);
INSERT INTO "llm_family_models" VALUES (15,'claude-v1.3-100k',3);
INSERT INTO "llm_family_models" VALUES (16,'claude-instant-1.1',3);