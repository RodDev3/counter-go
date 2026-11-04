export namespace main {
	
	export class Counter {
	    attempts: number;
	    success: number;
	    successPerfect: number;
	    currentStreak: number;
	    bestStreak: number;
	
	    static createFrom(source: any = {}) {
	        return new Counter(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.attempts = source["attempts"];
	        this.success = source["success"];
	        this.successPerfect = source["successPerfect"];
	        this.currentStreak = source["currentStreak"];
	        this.bestStreak = source["bestStreak"];
	    }
	}

}

