from django.contrib.auth.models import User
from farmers_markets.models import FarmersMarket
import csv


class Command(BaseCommand):
    help = "Import farmers markets from CSV file"

    def add_arguments(self, parser):
        parser.add_argument("csv_file", type=str)

    def handle(self, *args, **options):
        system_user = User.objects.get_or_create(username="system")[0]

        with open(options["csv_file"]) as file:
            reader = csv.DictReader(file)
            for row in reader:
                market = FarmersMarket(
                    name=row["name"],
                    operator=row["operator"] or "",
                    address=row["address"],
                    zip_code=row["zip"],
                    latitude=float(row["Y"]),
                    longitude=float(row["X"]),
                    contact_website=row["contact_website"] or "",
                    contact_phone=row["contact_phone"] or "",
                    contact_email=row["contact_email"] or "",
                    contact_facebook=row["contact_facebook"] or "",
                    contact_instagram=row["contact_instagram"] or "",
                    contact_twitter=row["contact_twitter"] or "",
                    # Combine hours and exceptions
                    hours_mon=self._format_hours(
                        row["hours_mon_start"],
                        row["hours_mon_end"],
                        row["hours_mon_exceptions"],
                    ),
                    hours_tue=self._format_hours(
                        row["hours_tues_start"],
                        row["hours_tues_end"],
                        row["hours_tues_exceptions"],
                    ),
                    # ... repeat for other days
                    season_year_round=row["season_year_round"].lower() == "yes",
                    season_opening_month=row["season_opening_month"] or "",
                    season_opening_day=int(row["season_opening_day"])
                    if row["season_opening_day"]
                    else None,
                    season_closing_month=row["season_closing_month"] or "",
                    season_closing_day=int(row["season_closing_day"])
                    if row["season_closing_day"]
                    else None,
                    accepts_credit=row["payment_credit"].lower() == "yes",
                    accepts_snap=row["payment_snap"].lower() == "yes",
                    accepts_fmnp=row["payment_fmnp"].lower() == "yes",
                    accepts_philly_food_bucks=row["payment_philly_food_bucks"].lower()
                    == "yes",
                    accepts_cash=row["payment_cash"].lower() == "yes",
                    payment_notes=row["payment_other_low_cost"] or "",
                    last_edited_by=system_user,
                )
                market.save()

                self.stdout.write(f"Imported: {market.name}")

    def _format_hours(self, start, end, exceptions):
        if not start or not end:
            return ""
        hours = f"{start}-{end}"
        if exceptions:
            hours += f" ({exceptions})"
        return hours

