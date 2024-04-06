# Tama Foundation Database Design & API Contract

users table

- id : bigint
- name : varchar
- occupation : varchar
- email : varchar
- password_hash : varchar
- token : varchar
- avatar_file_name : varchar
- role : varchar
- created_at
- updated_at

campaigns table

- id : bigint
- user_id : bigint
- name : varchar
- short_description : varchar
- description : text
- perks : text (comma separated)
- backer_count : integer
- goal_amount : integer
- current_amount : integer
- slug
- created_at
- updated_at

campaign_images table

- id : bigint
- campaign_id : bigint
- file_name : varchar
- is_primary : boolean

transactions table

- id : bigint
- campaign_id : integer
- user_id : bigint
- amount : integer
- status : varchar
- created_at
- updated_at

transaction_details table (?) ga butuh sih ini menurutku

## API Contract

**POST : api/v1/users**
**params :**
- name
- occupation
- email
- password

**response :**

```json
meta : {
	message: 'Your account has been created',
	code: 200,
	status: 'success'
},
data : {
	id: 1,
	name: "Tama",
	occupation: "content creator",
	email: "tama@gmail.com",
	token: "peterpanyangterdalam"
}
```

**POST : api/v1/email_checkers
params :**
- email

**response:**

```json
meta : {
	message: 'Email address has been registered',
	code: 200,
	status: 'success'
},
data : {
	is_available: false
}
```

**POST: api/v1/avatars
params:**
- avatar (form)

**response:**

```json
meta : {
	message: 'Avatar successfully uploaded,
	code: 200,
	status: 'success'
},
data : {
	is_uploaded: true
}
```

**POST: api/v1/sessions
params:**
- email
- password

**response:**

meta : {

message: 'You're now logged in'

code: 200

status: 'success'

},

data : {

id: 1,
name: "Tama",
occupation: "content creator",

email: "Tama@gmail.com",

token: "peterpanyangterdalam"

}

**GET: api/v1/campaigns
params:**
optional
- user_id
- backer_id
- none

**response:**

```json
meta : {
	message: 'List of campaigns',
	code: 200,
	status: 'success'
},
data : [
{
		id: 1,
		name: "BWA Startup",
		short_description: "Laris manis tanjung kimpul, mari belajar bareng",
		image_url: "domain/path/image.jpg",
		goal_amount: 1000000000,
		current_amount: 500000000,
		slug: "slug-here",
		user_id: 10
	}
]
```

**GET: api/v1/campaigns/1 
params:**
none

**response:**

```json
meta : {
	message: 'single campaigns',
	code: 200,
	status: 'success'
},
data : {
	id: 1,
	name: "BWA Startup",
	short_description: "Laris manis tanjung kimpul, mari belajar bareng",
	image_url: "path/image.jpg",
	goal_amount: 1000000000,
	current_amount: 500000000,
	user_id: 10,
  	slug: "slug",
	description: "Lorem epsum dolor sit amet yang 	panjang text-nya",
	user : {
		name: "Julia Ester",
		image_url: "path/image.jpg"
	},
	perks: [
		"Nintendo Switch",
		"Play Station 4"
	],
	images: [
		{
			image_url: "path/image.jpg",
			is_primary: true
		}
	]
}
```

POST: api/v1/campaigns

```json
{
    "meta": {
        "message": "Campaign successfully created ",
        "code": 200,
        "status": "success"
    },
    "data": {
        "user_id": 1,
        "goal_amount": 100000000,
		"current_amount": 0,
        "id": 7,
        "name": "Switch Pro",
        "image_url": "",
		"slug": "slug",
        "short_description": "Upcoming Nintendo Switch Pro"
    }
}
```

PUT : api/v1/campaigns/1

sama dengan atas

POST: api/v1/campaign-images
- file
- campaign_id
- is_primary

sama dengan upload avatar

GET : api/v1/campaigns/:id/transactions (campaign punya transaksi backer siapa aja)

```json
meta : {
	message: 'List of transactions'
	code: 200,
	status: 'success'
},
data : [
	{	
		id: 1,
		name: "Tama",
		amount: 1000000000,
		created_at: datetime
	}
]
```

GET : api/v1/transactions/ (user pernah transaksi apa aja)
params : header auth (current user)

```json
meta : {
	message: 'List of backed campaigns',
	code: 200,
	status: 'success'

},
data : [
	{	
		id: 1,
		amount: 1000000000,
		status: "paid",
		created_at: datetime,
		campaign: {
			name: "Hola",
			image_url : "path/to/file.png"
		}	
	}
]
```