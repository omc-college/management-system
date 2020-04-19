"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
exports.__esModule = true;
var core_1 = require("@angular/core");
var CalendarComponent = /** @class */ (function () {
    function CalendarComponent() {
        this.months = [
            'January',
            'February',
            'March',
            'April',
            'May',
            'June',
            'July',
            'August',
            'September',
            'October',
            'November',
            'December',
        ];
        this.days = []; // needed for rendering days in month
        this.beforeDays = []; // needed for rendering days before chosen month
        this.afterDays = []; // needed for rendering days after chosen month
        this.switcher = false; // switch between showing days or months
        this.chosenDate = new Date(); // date which is chosen at this moment
    }
    CalendarComponent.prototype.ngOnInit = function () {
        this.renderDays();
    };
    CalendarComponent.prototype.renderDays = function () {
        // fills days, beforeDays, AfterDays
        var daysInMonth = new Date(this.chosenDate.getFullYear(), this.chosenDate.getMonth() + 1, 0).getDate();
        // number of day in week of first day of month
        var startDay = new Date(this.chosenDate.getFullYear(), this.chosenDate.getMonth(), 1).getDay();
        var before = new Date(this.chosenDate.getFullYear(), this.chosenDate.getMonth(), 0).getDate() + 1;
        for (var i = startDay; i >= 0; i--) {
            this.beforeDays[startDay - i] = new Date(this.chosenDate.getFullYear(), this.chosenDate.getMonth() - 1, before - i);
        }
        // if we change year and previous date has more days than new date we just cut unuseful days
        this.beforeDays.splice(startDay);
        for (var day = 1; day <= daysInMonth; day++) {
            if (day === this.chosenDate.getDate()) {
                this.days[day - 1] = this.chosenDate;
                continue;
            }
            this.days[day - 1] = new Date(this.chosenDate.getFullYear(), this.chosenDate.getMonth(), day);
        }
        this.days.splice(daysInMonth);
        var after = 42 - daysInMonth - startDay;
        for (var i = 1; i <= after; i++) {
            this.afterDays[i - 1] = new Date(this.chosenDate.getFullYear(), this.chosenDate.getMonth() + 1, i);
        }
        this.afterDays.splice(after);
    };
    CalendarComponent.prototype.changeDay = function (e) {
        this.chosenDate = e;
    };
    CalendarComponent.prototype.changeDate = function (y, m) {
        this.chosenDate = new Date(y, m, this.chosenDate.getDate());
        this.renderDays();
    };
    CalendarComponent = __decorate([
        core_1.Component({
            selector: 'app-calendar',
            templateUrl: './calendar.component.html',
            styleUrls: ['./calendar.component.sass']
        })
    ], CalendarComponent);
    return CalendarComponent;
}());
exports.CalendarComponent = CalendarComponent;
