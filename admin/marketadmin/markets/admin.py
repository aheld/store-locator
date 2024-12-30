from django.contrib import admin

from .models import FarmersMarket, MarketHistory


class MarketHistoryInline(admin.StackedInline):
    model = MarketHistory
    extra = 0
    readonly_fields = ["edited_by", "edited_at"]
    can_delete = False

    def has_add_permission(self, request, obj=None):
        return False


@admin.register(FarmersMarket)
class FarmersMarketAdmin(admin.ModelAdmin):
    list_display = [
        "name",
        "address",
        "zip_code",
        "season_year_round",
        "updated_at",
        "last_edited_by",
    ]
    list_filter = ["season_year_round", "accepts_snap", "accepts_fmnp"]
    search_fields = ["name", "address"]
    inlines = [MarketHistoryInline]

    def save_model(self, request, obj, form, change):
        if change:
            # Create history record before saving changes
            MarketHistory.objects.create(
                market=obj,
                edited_by=request.user,
                name=obj.name,
                operator=obj.operator,
                address=obj.address,
                # ... copy other fields
            )

        obj.last_edited_by = request.user
        super().save_model(request, obj, form, change)
