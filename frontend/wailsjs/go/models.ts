export namespace helpers {
	
	export class WebsiteCredentials {
	    Name: string;
	    LoginInfo: {[key: string]: string};
	    Tokens: {[key: string]: };
	
	    static createFrom(source: any = {}) {
	        return new WebsiteCredentials(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.LoginInfo = source["LoginInfo"];
	        this.Tokens = source["Tokens"];
	    }
	}

}

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
	export class website {
	    Name: string;
	    Fields: string[];
	
	    static createFrom(source: any = {}) {
	        return new website(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Fields = source["Fields"];
	    }
	}

}

export namespace selfupdate {
	
	export class Release {
	    Version: semver.Version;
	    AssetURL: string;
	    AssetByteSize: number;
	    AssetID: number;
	    ValidationAssetID: number;
	    URL: string;
	    ReleaseNotes: string;
	    Name: string;
	    // Go type: time
	    PublishedAt?: any;
	    RepoOwner: string;
	    RepoName: string;
	
	    static createFrom(source: any = {}) {
	        return new Release(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Version = this.convertValues(source["Version"], semver.Version);
	        this.AssetURL = source["AssetURL"];
	        this.AssetByteSize = source["AssetByteSize"];
	        this.AssetID = source["AssetID"];
	        this.ValidationAssetID = source["ValidationAssetID"];
	        this.URL = source["URL"];
	        this.ReleaseNotes = source["ReleaseNotes"];
	        this.Name = source["Name"];
	        this.PublishedAt = this.convertValues(source["PublishedAt"], null);
	        this.RepoOwner = source["RepoOwner"];
	        this.RepoName = source["RepoName"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

export namespace semver {
	
	export class PRVersion {
	    VersionStr: string;
	    VersionNum: number;
	    IsNum: boolean;
	
	    static createFrom(source: any = {}) {
	        return new PRVersion(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.VersionStr = source["VersionStr"];
	        this.VersionNum = source["VersionNum"];
	        this.IsNum = source["IsNum"];
	    }
	}
	export class Version {
	    Major: number;
	    Minor: number;
	    Patch: number;
	    Pre: PRVersion[];
	    Build: string[];
	
	    static createFrom(source: any = {}) {
	        return new Version(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Major = source["Major"];
	        this.Minor = source["Minor"];
	        this.Patch = source["Patch"];
	        this.Pre = this.convertValues(source["Pre"], PRVersion);
	        this.Build = source["Build"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

export namespace variables {
	
	export class CompatibleDownloader {
	    Name: string;
	    Link: string;
	
	    static createFrom(source: any = {}) {
	        return new CompatibleDownloader(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Link = source["Link"];
	    }
	}
	export class General {
	    thumbnailsSize: number;
	    lang: string;
	    darkMode: boolean;
	    itemClickedAction: string;
	    xxx: boolean;
	
	    static createFrom(source: any = {}) {
	        return new General(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.thumbnailsSize = source["thumbnailsSize"];
	        this.lang = source["lang"];
	        this.darkMode = source["darkMode"];
	        this.itemClickedAction = source["itemClickedAction"];
	        this.xxx = source["xxx"];
	    }
	}
	export class Item {
	    Name: string;
	    Thumbnail: string;
	    Link: string;
	    Metadata: {[key: string]: string};
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Thumbnail = source["Thumbnail"];
	        this.Link = source["Link"];
	        this.Metadata = source["Metadata"];
	    }
	}
	export class ItemList {
	    Website: string;
	    CompatibleDownloaders: CompatibleDownloader[];
	    Items: Item[];
	
	    static createFrom(source: any = {}) {
	        return new ItemList(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Website = source["Website"];
	        this.CompatibleDownloaders = this.convertValues(source["CompatibleDownloaders"], CompatibleDownloader);
	        this.Items = this.convertValues(source["Items"], Item);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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
		    if (a.slice && a.map) {
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

