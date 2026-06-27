export namespace services {
	
	export class PasswordOptions {
	    length: number;
	    includeUpper: boolean;
	    includeLower: boolean;
	    includeDigits: boolean;
	    includeSpecial: boolean;
	    excludeAmbiguous: boolean;
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new PasswordOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.length = source["length"];
	        this.includeUpper = source["includeUpper"];
	        this.includeLower = source["includeLower"];
	        this.includeDigits = source["includeDigits"];
	        this.includeSpecial = source["includeSpecial"];
	        this.excludeAmbiguous = source["excludeAmbiguous"];
	        this.count = source["count"];
	    }
	}
	export class RemoveDuplicateOptions {
	    keepOrder: boolean;
	    ignoreCase: boolean;
	    removeEmpty: boolean;
	
	    static createFrom(source: any = {}) {
	        return new RemoveDuplicateOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.keepOrder = source["keepOrder"];
	        this.ignoreCase = source["ignoreCase"];
	        this.removeEmpty = source["removeEmpty"];
	    }
	}
	export class Result {
	    code: number;
	    message: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}
	export class SortLinesOptions {
	    order: string;
	    numeric: boolean;
	    ignoreCase: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SortLinesOptions(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.order = source["order"];
	        this.numeric = source["numeric"];
	        this.ignoreCase = source["ignoreCase"];
	    }
	}

}

