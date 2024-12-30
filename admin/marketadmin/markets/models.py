from django.db import models

from django.contrib.auth.models import User
from django.utils import timezone

from typing import Dict, Any


class FarmersMarket(models.Model):
    name = models.CharField(max_length=255)
    operator = models.CharField(max_length=255, blank=True)
    address = models.CharField(max_length=255)
    zip_code = models.CharField(max_length=10)
    latitude = models.DecimalField(max_digits=10, decimal_places=8)
    longitude = models.DecimalField(max_digits=11, decimal_places=8)
    contact_website = models.URLField(blank=True)
    contact_phone = models.CharField(max_length=50, blank=True)
    contact_email = models.EmailField(blank=True)
    contact_facebook = models.URLField(blank=True)
    contact_instagram = models.CharField(max_length=100, blank=True)
    contact_twitter = models.CharField(max_length=100, blank=True)

    # Operating hours stored as CharField to handle exceptions
    hours_mon = models.CharField(max_length=100, blank=True)
    hours_tue = models.CharField(max_length=100, blank=True)
    hours_wed = models.CharField(max_length=100, blank=True)
    hours_thu = models.CharField(max_length=100, blank=True)
    hours_fri = models.CharField(max_length=100, blank=True)
    hours_sat = models.CharField(max_length=100, blank=True)
    hours_sun = models.CharField(max_length=100, blank=True)

    season_year_round = models.BooleanField(default=False)
    season_opening_month = models.CharField(max_length=20, blank=True)
    season_opening_day = models.IntegerField(null=True, blank=True)
    season_closing_month = models.CharField(max_length=20, blank=True)
    season_closing_day = models.IntegerField(null=True, blank=True)

    accepts_credit = models.BooleanField(default=False)
    accepts_snap = models.BooleanField(default=False)
    accepts_fmnp = models.BooleanField(default=False)
    accepts_philly_food_bucks = models.BooleanField(default=False)
    accepts_cash = models.BooleanField(default=True)

    payment_notes = models.TextField(blank=True)

    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
    last_edited_by = models.ForeignKey(
        User, on_delete=models.SET_NULL, null=True, related_name="edited_markets"
    )

    def __str__(self):
        return str(self.name)

    class Meta:
        ordering = ["name"]


class MarketHistory(models.Model):
    market = models.ForeignKey(
        "FarmersMarket", on_delete=models.CASCADE, related_name="history"
    )
    edited_by = models.ForeignKey(User, on_delete=models.SET_NULL, null=True)
    edited_at = models.DateTimeField(default=timezone.now)
    data_snapshot: models.JSONField = (
        models.JSONField()
    )  # Stores complete market data at time of edit

    class Meta:
        ordering = ["-edited_at"]
        verbose_name_plural = "Market histories"

    def makeLinterHappy(self, jsontype) -> Dict[str, Any]:
        return jsontype

    def restore(self):
        data = self.makeLinterHappy(self.data_snapshot)
        market = self.market

        for field, value in data.items():
            if hasattr(market, field) and field not in ["id", "created_at"]:
                setattr(market, field, value)

        market.save()
