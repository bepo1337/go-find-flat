# Go-find-flat

This application is helping to find suitable flat advertisements. It automatically scrapes Kleinanzeigen every couple minutes with specified criteria and will send a Telegram message with relevant info and the link if new suitable advertisements have been uploaded.

To make this work you have to create a `.env` and include:
- Create a Telegram Bot and get the [API key](https://core.telegram.org/bots/tutorial#obtain-your-bot-token) `TELEGRAM_API_KEY=<YOUR_TELEGRAM_API_KEY`
- your chat-id `CHAT_ID=<YOUR_CHAT_ID>`

To customize searching behaviour, you have to modify the source code.
For example the search URL or the blacklisted postal codes. 

Right now ads are excluded if they:
- Have "Tauschwohnung" in their title
- Have "Untermiete" in their title
- Have a postal code in the blacklist

To prevent ads from being sent to Telegram twice, a JSON file will be updated to keep track of all ads we've seen so far.

Works for:
- [Kleinanzeigen](https://www.kleinanzeigen.de/) :ballot_box_with_check:
