CREATE TYPE role AS ENUM ('PG', 'SG', 'SF', 'PF', 'C');
CREATE TYPE conference AS ENUM ('EASTERN', 'WESTERN');

CREATE TABLE IN NOT EXISTS players (
    id uuid not null default gen_random_uuid(),
    name varchar(256) not null,
    surname varchar(256) not null,
    height smallint not null CHECK (height > 100 AND height < 310),
    weight smallint not null CHECK (weight > 20 AND weight < 400),
    citizenship varchar(256) not null,
    role role not null,
    team_id uuid not null references teams(id),
    primary_key(id)
)


CREATE TABLE IN NOT EXISTS teams (
    id uuid not null default gen_random_uuid(),
    name varchar(256) not null,
    conf conference not null,
    primary_key(id)
)


CREATE TABLE IN NOT EXISTS games (
    id uuid not null default gen_random_uuid(),
    home_team_id uuid not null references teams(id),
    guest_team_id uuid not null references teams(id),
    game_date date not null,
    home_score smallint not null,
    guest_score smallint not null,
    primary_key(id)
)

CREATE TABLE IN NOT EXISTS statistics (
    player_id uuid not null references players(id),
    game_id uuid not null references games(id),
    points smallint not null,
    assists smallint not null,
    rebound smallint not null,
    interceptions smallint not null,
    primary_key(player_id, game_id)
)
