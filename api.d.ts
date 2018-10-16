declare enum NodeType {
    ErrorNode = 0,
    TextNode = 1,
    DocumentNode = 2,
    ElementNode = 3,
    CommentNode = 4,
    DoctypeNode = 5
}

declare enum Atom {
    a = 0x1,
    abbr = 0x4,
    accept = 0x1a06,
    acceptCharset = 0x1a0e,
    accesskey = 0x2c09,
    acronym = 0xaa07,
    action = 0x27206,
    address = 0x6f307,
    align = 0xb105,
    allowfullscreen = 0x2080f,
    allowpaymentrequest = 0xc113,
    allowusermedia = 0xdd0e,
    alt = 0xf303,
    annotation = 0x1c90a,
    annotationXml = 0x1c90e,
    applet = 0x31906,
    area = 0x35604,
    article = 0x3fc07,
    as = 0x3c02,
    aside = 0x10705,
    async = 0xff05,
    audio = 0x11505,
    autocomplete = 0x2780c,
    autofocus = 0x12109,
    autoplay = 0x13c08,
    b = 0x101,
    base = 0x3b04,
    basefont = 0x3b08,
    bdi = 0xba03,
    bdo = 0x14b03,
    bgsound = 0x15e07,
    big = 0x17003,
    blink = 0x17305,
    blockquote = 0x1870a,
    body = 0x2804,
    br = 0x202,
    button = 0x19106,
    canvas = 0x10306,
    caption = 0x23107,
    center = 0x22006,
    challenge = 0x29b09,
    charset = 0x2107,
    checked = 0x47907,
    cite = 0x19c04,
    class = 0x56405,
    code = 0x5c504,
    col = 0x1ab03,
    colgroup = 0x1ab08,
    color = 0x1bf05,
    cols = 0x1c404,
    colspan = 0x1c407,
    command = 0x1d707,
    content = 0x58b07,
    contenteditable = 0x58b0f,
    contextmenu = 0x3800b,
    controls = 0x1de08,
    coords = 0x1ea06,
    crossorigin = 0x1fb0b,
    data = 0x4a504,
    datalist = 0x4a508,
    datetime = 0x2b808,
    dd = 0x2d702,
    default = 0x10a07,
    defer = 0x5c705,
    del = 0x45203,
    desc = 0x56104,
    details = 0x7207,
    dfn = 0x8703,
    dialog = 0xbb06,
    dir = 0x9303,
    dirname = 0x9307,
    disabled = 0x16408,
    div = 0x16b03,
    dl = 0x5e602,
    download = 0x46308,
    draggable = 0x17a09,
    dropzone = 0x40508,
    dt = 0x64b02,
    em = 0x6e02,
    embed = 0x6e05,
    enctype = 0x28d07,
    face = 0x21e04,
    fieldset = 0x22608,
    figcaption = 0x22e0a,
    figure = 0x24806,
    font = 0x3f04,
    footer = 0xf606,
    for = 0x25403,
    foreignObject = 0x2540d,
    foreignobject = 0x2610d,
    form = 0x26e04,
    formaction = 0x26e0a,
    formenctype = 0x2890b,
    formmethod = 0x2a40a,
    formnovalidate = 0x2ae0e,
    formtarget = 0x2c00a,
    frame = 0x8b05,
    frameset = 0x8b08,
    h1 = 0x15c02,
    h2 = 0x2de02,
    h3 = 0x30d02,
    h4 = 0x34502,
    h5 = 0x34f02,
    h6 = 0x64d02,
    head = 0x33104,
    header = 0x33106,
    headers = 0x33107,
    height = 0x5206,
    hgroup = 0x2ca06,
    hidden = 0x2d506,
    high = 0x2db04,
    hr = 0x15702,
    href = 0x2e004,
    hreflang = 0x2e008,
    html = 0x5604,
    httpEquiv = 0x2e80a,
    i = 0x601,
    icon = 0x58a04,
    id = 0x10902,
    iframe = 0x2fc06,
    image = 0x30205,
    img = 0x30703,
    input = 0x44b05,
    inputmode = 0x44b09,
    ins = 0x20403,
    integrity = 0x23f09,
    is = 0x16502,
    isindex = 0x30f07,
    ismap = 0x31605,
    itemid = 0x38b06,
    itemprop = 0x19d08,
    itemref = 0x3cd07,
    itemscope = 0x67109,
    itemtype = 0x31f08,
    kbd = 0xb903,
    keygen = 0x3206,
    keytype = 0xd607,
    kind = 0x17704,
    label = 0x5905,
    lang = 0x2e404,
    legend = 0x18106,
    li = 0xb202,
    link = 0x17404,
    list = 0x4a904,
    listing = 0x4a907,
    loop = 0x5d04,
    low = 0xc303,
    main = 0x1004,
    malignmark = 0xb00a,
    manifest = 0x6d708,
    map = 0x31803,
    mark = 0xb604,
    marquee = 0x32707,
    math = 0x32e04,
    max = 0x33d03,
    maxlength = 0x33d09,
    media = 0xe605,
    mediagroup = 0xe60a,
    menu = 0x38704,
    menuitem = 0x38708,
    meta = 0x4b804,
    meter = 0x9805,
    method = 0x2a806,
    mglyph = 0x30806,
    mi = 0x34702,
    min = 0x34703,
    minlength = 0x34709,
    mn = 0x2b102,
    mo = 0xa402,
    ms = 0x67402,
    mtext = 0x35105,
    multiple = 0x35f08,
    muted = 0x36705,
    name = 0x9604,
    nav = 0x1303,
    nobr = 0x3704,
    noembed = 0x6c07,
    noframes = 0x8908,
    nomodule = 0xa208,
    nonce = 0x1a605,
    noscript = 0x21608,
    novalidate = 0x2b20a,
    object = 0x26806,
    ol = 0x13702,
    onabort = 0x19507,
    onafterprint = 0x2360c,
    onautocomplete = 0x2760e,
    onautocompleteerror = 0x27613,
    onauxclick = 0x61f0a,
    onbeforeprint = 0x69e0d,
    onbeforeunload = 0x6e70e,
    onblur = 0x56d06,
    oncancel = 0x11908,
    oncanplay = 0x14d09,
    oncanplaythrough = 0x14d10,
    onchange = 0x41b08,
    onclick = 0x2f507,
    onclose = 0x36c07,
    oncontextmenu = 0x37e0d,
    oncopy = 0x39106,
    oncuechange = 0x3970b,
    oncut = 0x3a205,
    ondblclick = 0x3a70a,
    ondrag = 0x3b106,
    ondragend = 0x3b109,
    ondragenter = 0x3ba0b,
    ondragexit = 0x3c50a,
    ondragleave = 0x3df0b,
    ondragover = 0x3ea0a,
    ondragstart = 0x3f40b,
    ondrop = 0x40306,
    ondurationchange = 0x41310,
    onemptied = 0x40a09,
    onended = 0x42307,
    onerror = 0x42a07,
    onfocus = 0x43107,
    onhashchange = 0x43d0c,
    oninput = 0x44907,
    oninvalid = 0x45509,
    onkeydown = 0x45e09,
    onkeypress = 0x46b0a,
    onkeyup = 0x48007,
    onlanguagechange = 0x48d10,
    onload = 0x49d06,
    onloadeddata = 0x49d0c,
    onloadedmetadata = 0x4b010,
    onloadend = 0x4c609,
    onloadstart = 0x4cf0b,
    onmessage = 0x4da09,
    onmessageerror = 0x4da0e,
    onmousedown = 0x4e80b,
    onmouseenter = 0x4f30c,
    onmouseleave = 0x4ff0c,
    onmousemove = 0x50b0b,
    onmouseout = 0x5160a,
    onmouseover = 0x5230b,
    onmouseup = 0x52e09,
    onmousewheel = 0x53c0c,
    onoffline = 0x54809,
    ononline = 0x55108,
    onpagehide = 0x5590a,
    onpageshow = 0x5730a,
    onpaste = 0x57f07,
    onpause = 0x59a07,
    onplay = 0x5a406,
    onplaying = 0x5a409,
    onpopstate = 0x5ad0a,
    onprogress = 0x5b70a,
    onratechange = 0x5cc0c,
    onrejectionhandled = 0x5d812,
    onreset = 0x5ea07,
    onresize = 0x5f108,
    onscroll = 0x60008,
    onsecuritypolicyviolation = 0x60819,
    onseeked = 0x62908,
    onseeking = 0x63109,
    onselect = 0x63a08,
    onshow = 0x64406,
    onsort = 0x64f06,
    onstalled = 0x65909,
    onstorage = 0x66209,
    onsubmit = 0x66b08,
    onsuspend = 0x67b09,
    ontimeupdate = 0x400c,
    ontoggle = 0x68408,
    onunhandledrejection = 0x68c14,
    onunload = 0x6ab08,
    onvolumechange = 0x6b30e,
    onwaiting = 0x6c109,
    onwheel = 0x6ca07,
    open = 0x1a304,
    optgroup = 0x5f08,
    optimum = 0x6d107,
    option = 0x6e306,
    output = 0x51d06,
    p = 0xc01,
    param = 0xc05,
    pattern = 0x6607,
    picture = 0x7b07,
    ping = 0xef04,
    placeholder = 0x1310b,
    plaintext = 0x1b209,
    playsinline = 0x1400b,
    poster = 0x2cf06,
    pre = 0x47003,
    preload = 0x48607,
    progress = 0x5b908,
    prompt = 0x53606,
    public = 0x58606,
    q = 0xcf01,
    radiogroup = 0x30a,
    rb = 0x3a02,
    readonly = 0x35708,
    referrerpolicy = 0x3d10e,
    rel = 0x48703,
    required = 0x24c08,
    reversed = 0x8008,
    rows = 0x9c04,
    rowspan = 0x9c07,
    rp = 0x23c02,
    rt = 0x19a02,
    rtc = 0x19a03,
    ruby = 0xfb04,
    s = 0x2501,
    samp = 0x7804,
    sandbox = 0x12907,
    scope = 0x67505,
    scoped = 0x67506,
    script = 0x21806,
    seamless = 0x37108,
    section = 0x56807,
    select = 0x63c06,
    selected = 0x63c08,
    shape = 0x1e505,
    size = 0x5f504,
    sizes = 0x5f505,
    slot = 0x1ef04,
    small = 0x20605,
    sortable = 0x65108,
    sorted = 0x33706,
    source = 0x37806,
    spacer = 0x43706,
    span = 0x9f04,
    spellcheck = 0x4740a,
    src = 0x5c003,
    srcdoc = 0x5c006,
    srclang = 0x5f907,
    srcset = 0x6f906,
    start = 0x3fa05,
    step = 0x58304,
    strike = 0xd206,
    strong = 0x6dd06,
    style = 0x6ff05,
    sub = 0x66d03,
    summary = 0x70407,
    sup = 0x70b03,
    svg = 0x70e03,
    system = 0x71106,
    tabindex = 0x4be08,
    table = 0x59505,
    target = 0x2c406,
    tbody = 0x2705,
    td = 0x9202,
    template = 0x71408,
    textarea = 0x35208,
    tfoot = 0xf505,
    th = 0x15602,
    thead = 0x33005,
    time = 0x4204,
    title = 0x11005,
    tr = 0xcc02,
    track = 0x1ba05,
    translate = 0x1f209,
    tt = 0x6802,
    type = 0xd904,
    typemustmatch = 0x2900d,
    u = 0xb01,
    ul = 0xa702,
    updateviacache = 0x460e,
    usemap = 0x59e06,
    value = 0x1505,
    var = 0x16d03,
    video = 0x2f105,
    wbr = 0x57c03,
    width = 0x64905,
    workertype = 0x71c0a,
    wrap = 0x72604,
    xmp = 0x12f03
}

declare type Sanitizer = (node: Node, parent: Node) => Error;
declare type Sanitization = [String, Sanitizer];

declare type Extractor = (node: Node) => String;
declare type Predicate = (node: Node) => boolean;

declare interface String {
    replaceAll(oldValue: String, newValue: String, repeats: number): String;

    startsWith(value: String): boolean;

    endsWith(value: String): boolean;
}

declare interface Node {
    Parent: Node;
    FirstChild: Node;
    LastChild: Node;
    PrevSibling: Node;
    NextSibling: Node;
    Type: NodeType;
    DataAtom: Atom;
    Data: String;
    Namespace: String;
    Attr: Attribute[];
}

declare class Style {
    getDeclaration(key: String): [String, boolean];

    setDeclaration(key: String, value: String): void;

    removeDeclaration(key: String): void;

    computeStyle(): String;

    attachStyle(node: Node): void;
}

declare interface Attribute {
    Namespace: String;
    Key: String;
    Val: String;
}

interface Api {
    isTextOnly(node: Node): boolean;

    parseStyle(node: Node): Style;

    hasAttribute(node: Node, attribute: String): boolean;

    getAttribute(node: Node, attribute: String): [Attribute, boolean];

    setAttribute(node: Node, attribute: String, value: String): void;

    removeAttribute(node: Node, attribute: String): void;

    createTag(tagName: String): Node;

    createTextNode(value: String): Node;

    moveNode(node: Node, oldParent: Node, newParent: Node): void;

    removeNode(parent: Node, node: Node): void;

    emptyNode(node: Node): void;

    replaceNode(parent: Node, newNode: Node, oldNode: Node): void;

    children(node: Node): Node[];

    cloneNode(node: Node): Node;

    appendNode(node: Node, parent: Node): void;
}


declare const api: Api;

declare function DeleteNodeAndChildren(node: Node, parent: Node): Error;
declare function DeleteElementAndMoveChildrenToParent(node: Node, parent: Node): Error;
declare function ReplaceElementAndReassignChildren(tagName: String): Sanitizer;
declare function InjectOuterElement(tagName: String): Sanitizer;

declare function SetStyleDeclaration(property: String, value: String): Sanitizer;
declare function DeleteStyleDeclaration(property: String): Sanitizer;

declare function SelectParent(sanitizer: Sanitizer): Sanitizer;
declare function And(...sanitizers: Sanitizer[]): Sanitizer;

declare function SetAttribute(attribute: String, value: String): Sanitizer;
declare function SetAttributeWithExtractor(attribute: String, extractor: Extractor): Sanitizer;
declare function DeleteAttribute(attribute: String): Sanitizer;

declare function Filter(predicate: Predicate, ...sanitizers: Sanitizer[]) : Sanitizer;
declare function Filters(...predicates: Predicate[]) : Predicate;

declare function sanitize(selector: String, sanitizer: Sanitizer): Sanitization;

declare function sanitizers(...sanitizations: Sanitization[]): void;
