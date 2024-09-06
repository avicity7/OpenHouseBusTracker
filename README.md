![image](https://i.imgur.com/4Z8PMqQ.png)

## What we do

Through this platform, we digitise the operations needed to manage and track the Buses used for Singapore Polytechnic's Open House event in realtime.


## Tech Stack

**Frontend:** SvelteKit, TailwindCSS

**Backend:** Go, PostgreSQL

## Installation
Run `cd frontend` then `yarn` to install dependencies. Get the frontend started by running `yarn dev`.

From the `/frontend` folder, run `cd ../backend` to get into the backend folder and run `go run .`. Installation of dependencies will be automatically handled.

## Environment Variables

You'll need .env variables to run everything, contact [@avicity7](https://github.com/avicity7) if you don't have them.

### Frontend Environment Variables

Create a `.env` file in the `frontend` directory with the following variables:

**Example `.env` for Frontend:**
```env
PUBLIC_BACKEND_URL=http://127.0.0.1
PUBLIC_MAPBOX_KEY=your_mapbox_key_here
PUBLIC_OSR_KEY=your_osr_key_here
```

### Backend Environment Variables

Create a `.env` file in the `backend` directory with the following variables:

**Example `.env` for Backend:**
```env
DATABASE_URL=your_database_url_here
SECRET=your_secret_key_here
ENV=DEV
EMAIL_PASSWORD=your_email_password_here
```

## Deployment

We're live at [https://open-house-bus-tracker.vercel.app/]!

