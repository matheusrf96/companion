// c = window
// o = document
// m = 'HASH'

(function(c,o,m,r,a,d,e){
c.cmp=c.cmp||function(){(c.cmp.q=c.cmp.q||[]).push(arguments)};  // Set cmp query as request arguments
r=o.createElement('input');r.type='hidden';r.id='eh';r.value=m;  // Create the following input `<input type="hidden" id="eh" value="{m}" />`
d=o.createElement('script');d.async=1;d.src='http://localhost:8000/cmp';  // Create the following script `<script src="http://localhost:8000/cmp"></script>`
e=o.getElementsByTagName('body')[0];e.appendChild(r);  // Insert input "eh" into body
a=o.getElementsByTagName('script')[0];a.parentNode.insertBefore(d, a);  // Insert created script into DOM before this script
})(window,document,'<HASH>');

