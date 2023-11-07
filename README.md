# FindHotel Coding Challenge

This project allows to import and access geolocation data over HTTP. The project has two main components:
1. `viodata` - library for importing, validating and accessing saved data(in the Database).
2. `vioapi` - application for using `viodata` as importer or as a datasource.

## How to run?
Ensure you've setup correctly `.env` environment variables. You can make it by copying `.env.example` file.

For running project you need to install docker-compose. Link: https://docs.docker.com/compose/install/

Once the docker-compose is ready you can execute following command in the root folder `vio`:
```
docker-compose up
```

This command will setup `importer` and `api` services as well as dependent services like: `postgreSQL`, `pgWeb` and `migration`.

Almost all components related to have ready `postgreSQL` first. Once it's ready `migrating` process will start to ensure the schema of the database is correct. You can open `pgweb` page by following link: `https://localhost:8081` and check the schema.

Importing process will start in parallel with serving api. Entire process of importing can take a while depens on your docker configuration(like vCPU, vRAM, ect.) and local configuration of importing. You can find this configuration in the `.env` file. 

There are two main configurations:
1. `COUNT_GO_ROUTINE` - define amount of concurrent goroutine can handle(validate and insert) each row from imported file. 
2. `COUNT_BULK_INSERT` - define maximum amount of `ip_location` entities can be inserted in bulk operation to the database. It might be less entities in case when we rich the end of the imported file.

And one additional parameter in the docker-compose `max_connections=1000` for database describes maximum count of connections. By default in equal `100`, but has changed to `1000` to be more flexible with main configs. Exceeding the stated number will result in an error stating "too many clients already".

### Expected result:
With example local configuration from `.env.example` and setup: CPUs - 8, Memory - 16 GB(MacBook Pro 16 M1 PRO) importing result is:
```
{"importing_duration": "2m49.577327452s", "accepted_entries": 950076, "discarded_entries": 49924}
```

To access API follow the link: `https://localhost:8080/` it will automatically redirects to swagger documentation. At the same page located API query to get `ip_location` data by providing `IP Address`. By hitting `Try it out` the HTTP Request can be executed with custom `IP Address`: 
![API Example](api_example.png)

## TODO:

- [ ] API  | Introduce pagination, if there are more then one result on getting `ip_location` by `IP Address`;
- [ ] DATA | Improve storing `country_code` and `city` fields by including `country_code` to `countries` table and `city` as additional table linked to `country` and `ip_locations`. Right now it's not implemented due to slowing down process of importing and incorrect data(for example country `Netherlands` can have different country code like `PA`, `MT`);
- [ ] DATA | Store geolocation data in geometry Point object(GIS) instead of two columns. It will simplify GEO-queries in the future. Right now database already supporting GIS extension and there is a query for getting and saving geometry object. But, it will slow down importing due to limitation of using `copyfrom` feature.

### Notes:
- Database normalization will slow down speed of importing. But it can be improved with additional job that will take care of creating additional keys when it's required.