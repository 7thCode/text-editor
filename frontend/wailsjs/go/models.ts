export namespace main {
	
	export class FileData {
	    content: string;
	    path: string;
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new FileData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.content = source["content"];
	        this.path = source["path"];
	        this.name = source["name"];
	    }
	}

}

