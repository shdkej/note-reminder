mkdir build
pip install -t build -r recommender/requirments.txt
cd build
rm -r *.dist-info __pycache__
find . -name "__pycache__" -type d -exec rm -rf {} +
rm -rf numpy* bin dateutil threadpoolctl.py six.py scipy*
zip -r ../event_handling/recommand.zip .
cd ../event_handling
zip -r recommend.zip recommender_system.py
cd ../
