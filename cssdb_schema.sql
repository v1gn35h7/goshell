CREATE KEYSPACE goshell WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'}  AND durable_writes = true;

CREATE TABLE goshell.trooper_events (
    id text PRIMARY KEY,
    name text
);

 CREATE TABLE goshell.scripts (
    id text PRIMARY KEY,
    title text,
    script text,
    platform text,
    type text
);


 CREATE TABLE goshell.assets (
    AgentId text PRIMARY KEY,
    HostName text,
    architecture text,
    platform text,
    OperatingSystem text,
    SyncTime text
);