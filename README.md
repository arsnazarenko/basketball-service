# Basketball service


## API

### Players
* GET       api/v1/players
* GET       api/v1/players/{rd}
* POST      api/v1/players
* PUT       api/v1/players/{id}
* DELETE    api/v1/players/{id}

### Teams
* GET       api/v1/teams
* GET       api/v1/teams/{id}
* POST      api/v1/teams/
* PUT       api/v1/teams/{id}
* DELETE    api/v1/teams/{id}

* GET       api/v1/teams/{id}/players -- Non CRUD Method

### Games
* GET       api/v1/games
* GET       api/v1/games/{id}
* POST      api/v1/games/
* PUT       api/v1/games/{id}
* DELETE    api/v1/games/{id}

### Statistics
* GET       api/v1/statistics
* GET       api/v1/statistics/{id}
* POST      api/v1/statistics/
* PUT       api/v1/statistics/{id}
* DELETE    api/v1/statistics/{id}

* GET       api/v1/statistics/{player_id} -- Non CRUD Method
* GET       api/v1/statistics/{game_id}   -- Non CRUD Method
## DB Schema

## Entities
### Player
* ID
* Name
* Surname
* Age
* Height
* Weight
* Citizenship
* Role
* TeamID

### Team (One To Many: Commands -> Players)
* ID
* Name
* Conference (Sounth, North, West, East)

### Game 
* ID
* HomeTeamID
* GuestTeamID
* Type
* Date
* Winner (bool or one byte type with 0 or 1 values)

### Statistics (Many To Many: Players -> Statistics <- Games)
* ID
* PlayerID
* GameID
* Points
* Assists
* Rebounds
* Interceptions
