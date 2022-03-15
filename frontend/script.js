// c = window
// o = document
// m = 'HASH'
// r = Handler script creation
// a = Document "script" tag
// d = Hidden input with ecommerce hash creation
// e = Document "body" tag

(function(c,o,m,r,a,d,e){
r=o.createElement('script');r.async=1;r.src='http://localhost:8000/cmp'+c.location.search;  // Create script calling the handler with URL params
a=o.getElementsByTagName('script')[0];a.parentNode.insertBefore(r, a);  // Insert handler js before this script
d=o.createElement('input');d.type='hidden';d.id='eh';d.value=m  // Create a hidden input with ecommerce hash value
e=o.getElementsByTagName('body')[0];e.appendChild(d);  // Insert ecommerce hash input into body
})(window,document,'<HASH>');
