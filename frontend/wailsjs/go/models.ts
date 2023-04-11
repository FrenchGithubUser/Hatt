export namespace main {
	
	export class CustomList {
	    name: string;
	    sources: string[];
	
	    static createFrom(source: any = {}) {
	        return new CustomList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.sources = source["sources"];
	    }
	}
	export class General {
	    thumbnailsSize: number;
	    lang: string;
	    darkMode: boolean;
	    itemClickedAction: string;
	
	    static createFrom(source: any = {}) {
	        return new General(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.thumbnailsSize = source["thumbnailsSize"];
	        this.lang = source["lang"];
	        this.darkMode = source["darkMode"];
	        this.itemClickedAction = source["itemClickedAction"];
	    }
	}
	export class UserSettings {
	    general: General;
	
	    static createFrom(source: any = {}) {
	        return new UserSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.general = this.convertValues(source["general"], General);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

