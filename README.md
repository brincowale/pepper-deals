# Pepper Deals

## Overview

Pepper Deals is a project designed to fetch and filter deals from the Pepper API, sending notifications directly to a Telegram channel. It utilizes a SQLite database to store deals, and all configurations are managed through a `config.json` file.

Pepper aggregates deals from various websites, including mydealz.de, hotukdeals.co.uk, chollometro, and others, providing users with a comprehensive view of the best offers available.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- A Telegram channel to receive notifications.
- Access to the Telegram API by creating a bot using [@BotFather](https://t.me/botfather).
- A valid configuration file (`config.json`).

## Configuration

The application requires a configuration file named `config.json`, which should be placed in the root directory of the project. You can use the example provided in the repository as a template, which includes the necessary keys and secrets for mydealz.de.

### Filters

The filters are defined using regular expressions (regex), allowing for flexible and powerful pattern matching. Each filter object can contain the following fields:

- **include**: A regex pattern to include items that match the criteria.
- **exclude**: A regex pattern to exclude items that match the criteria.
- **include_website**: A regex pattern to include specific websites.
- **exclude_website**: A regex pattern to exclude specific websites.
- **lowest_price**: The minimum price for filtering.
- **maximum_price**: The maximum price for filtering.

### Example Configuration Structure

Below is the structure of the `config.json` file:

```json
{
  "telegram_api_key": "YOUR_TELEGRAM_API_KEY",
  "telegram_channel": "YOUR_TELEGRAM_CHANNEL",
  "consumer_key": "CONSUMER_KEY_FROM_LOCAL_PEPPER_SITE",
  "consumer_secret": "CONSUMER_SECRET_FROM_LOCAL_PEPPER_SITE",
  "host": "PEPPER_HOST",
  "pkgname": "APK_PACKAGE_NAME",
  "filters": [
    {
      "include": "INCLUDE_PATTERN",
      "exclude": "EXCLUDE_PATTERN",
      "include_website": "INCLUDE_WEBSITE_PATTERN",
      "exclude_website": "EXCLUDE_WEBSITE_PATTERN",
      "lowest_price": 0.0,
      "maximum_price": 100.0
    }
  ]
}```
